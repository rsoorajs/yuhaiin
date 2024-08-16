package tun2socket

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand/v2"
	"net"

	"github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/netlink"
	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/tun/device"
	"github.com/Asutorufa/yuhaiin/pkg/utils/pool"
	"gvisor.dev/gvisor/pkg/tcpip/header"
	i4 "gvisor.dev/gvisor/pkg/tcpip/network/ipv4"
	i6 "gvisor.dev/gvisor/pkg/tcpip/network/ipv6"
)

type UDP struct {
	device       netlink.Tun
	HandlePacket func(tuple Tuple, payload []byte)
	closed       bool
}

func NewUDP(device netlink.Tun) *UDP {
	return &UDP{device: device, HandlePacket: func(tuple Tuple, payload []byte) {}}
}

func (u *UDP) Close() error {
	u.closed = true
	return nil
}

func (u *UDP) handleUDPPacket(tuple Tuple, payload []byte) {
	if u.closed {
		return
	}
	u.HandlePacket(tuple, payload)
}

func (u *UDP) WriteTo(buf []byte, tuple Tuple) (int, error) {
	if u.closed {
		return 0, net.ErrClosed
	}

	tunBuf, err := u.processUDPPacket(buf, tuple)
	if err != nil {
		return 0, err
	}
	defer pool.PutBytes(tunBuf)

	_, err = u.device.Write([][]byte{tunBuf})
	return len(buf), err
}

type Batch struct {
	Payload []byte
	Tuple   Tuple
}

func (u *UDP) WriteBatch(batch []Batch) error {
	if u.closed {
		return net.ErrClosed
	}

	buffs := make([][]byte, 0, len(batch))

	for _, b := range batch {
		tunBuf, err := u.processUDPPacket(b.Payload, b.Tuple)
		if err != nil {
			log.Error("process udp packet failed:", "err", err)
			continue
		}
		defer pool.PutBytes(tunBuf)

		buffs = append(buffs, tunBuf)
	}

	if len(buffs) == 0 {
		return nil
	}

	_, err := u.device.Write(buffs)
	return err
}

func (u *UDP) processUDPPacket(buf []byte, tuple Tuple) ([]byte, error) {
	udpTotalLength := int(header.UDPMinimumSize) + len(buf)

	if udpTotalLength > math.MaxUint16 || udpTotalLength > int(u.device.MTU()) { // ip packet max length
		return nil, fmt.Errorf("udp packet too large: %d", len(buf))
	}

	tunBuf := pool.GetBytes(u.device.MTU() + u.device.Offset())

	ipBuf := tunBuf[u.device.Offset():]

	var ip header.Network
	var totalLength uint16

	if tuple.SourceAddr.Value().Len() == 4 && !tuple.DestinationAddr.Value().To4().Unspecified() {
		if tuple.DestinationAddr.Value().To4().Unspecified() {
			// return 0, fmt.Errorf("send IPv6 packet to IPv4 connection: src: %v, dst: %v", tuple.SourceAddr, tuple.DestinationAddr)
			slog.Warn("send IPv6 packet to IPv4 connection", slog.String("src", tuple.SourceAddr.Value().String()), slog.String("dst", tuple.DestinationAddr.Value().String()))
		}

		// no ipv4 options set, so ipv4 header size is IPv4MinimumSize
		totalLength = header.IPv4MinimumSize + uint16(udpTotalLength)

		ipv4 := header.IPv4(ipBuf)
		ipv4.Encode(&header.IPv4Fields{
			TOS:            0,
			ID:             uint16(rand.Uint32()),
			TotalLength:    totalLength,
			FragmentOffset: 0,
			TTL:            i4.DefaultTTL,
			Protocol:       uint8(header.UDPProtocolNumber),
			SrcAddr:        tuple.DestinationAddr.Value(),
			DstAddr:        tuple.SourceAddr.Value(),
		})

		ip = ipv4
	} else {
		// ipv6 header size is fixed
		totalLength = header.IPv6FixedHeaderSize + uint16(udpTotalLength)

		ipv6 := header.IPv6(ipBuf)
		ipv6.Encode(&header.IPv6Fields{
			TransportProtocol: header.UDPProtocolNumber,
			PayloadLength:     uint16(udpTotalLength),
			SrcAddr:           tuple.DestinationAddr.Value(),
			DstAddr:           tuple.SourceAddr.Value(),
			HopLimit:          i6.DefaultTTL,
			TrafficClass:      0,
		})

		ip = ipv6
	}

	udp := header.UDP(ip.Payload())

	udp.Encode(&header.UDPFields{
		SrcPort: tuple.DestinationPort,
		DstPort: tuple.SourcePort,
		Length:  uint16(udpTotalLength),
	})
	copy(udp.Payload(), buf)

	device.ResetIPChecksum(ip)

	// On IPv4, UDP checksum is optional, and a zero value indicates the
	// transmitter skipped the checksum generation (RFC768).
	// On IPv6, UDP checksum is not optional (RFC2460 Section 8.1).
	if _, ok := ip.(header.IPv6); ok {
		pseudoSum := header.PseudoHeaderChecksum(header.UDPProtocolNumber,
			ip.SourceAddress(), ip.DestinationAddress(), uint16(len(ip.Payload())))
		device.ResetTransportChecksum(ip, udp, pseudoSum)
	}

	return tunBuf[:totalLength+uint16(u.device.Offset())], nil
}
