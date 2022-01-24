package shadowsocksr

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/proxy"
	streamCipher "github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/cipher"
	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/obfs"
	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/protocol"
	ssr "github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/utils"
	socks5client "github.com/Asutorufa/yuhaiin/pkg/net/proxy/socks5/client"
)

var _ proxy.Proxy = (*Shadowsocksr)(nil)

type Shadowsocksr struct {
	encryptMethod   string
	encryptPassword string
	obfs            string
	obfsParam       string
	obfsData        interface{}
	protocol        string
	protocolParam   string
	protocolData    interface{}

	cipher *streamCipher.Cipher
	p      proxy.Proxy
}

func NewShadowsocksr(method, password, obfs, obfsParam, protocol, protocolParam string) func(proxy.Proxy) (proxy.Proxy, error) {
	return func(p proxy.Proxy) (proxy.Proxy, error) {
		cipher, err := streamCipher.NewCipher(method, password)
		if err != nil {
			return nil, err
		}

		s := &Shadowsocksr{
			encryptMethod:   method,
			encryptPassword: password,
			obfs:            obfs,
			obfsParam:       obfsParam,
			protocol:        protocol,
			protocolParam:   protocolParam,

			cipher: cipher,
			p:      p,
		}
		// s.protocolData = new(Protocol.AuthData)
		return s, nil
	}
}

func (s *Shadowsocksr) Conn(addr string) (net.Conn, error) {
	c, err := s.p.Conn(addr)
	if err != nil {
		return nil, fmt.Errorf("get conn failed: %w", err)
	}

	// should initialize obfs/protocol now
	rs, portS, err := net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		return nil, err
	}
	port, err := strconv.ParseUint(portS, 10, 16)
	if err != nil {
		return nil, err
	}

	obfs := obfs.NewObfs(c, s.obfs)
	if obfs == nil {
		return nil, errors.New("[ssr] unsupported obfs type: " + s.obfs)
	}
	protocol := protocol.NewProtocol(s.protocol)
	if protocol == nil {
		return nil, errors.New("[ssr] unsupported protocol type: " + s.protocol)
	}

	overhead := obfs.GetOverhead() + protocol.GetOverhead()

	obfs.SetServerInfo(&ssr.ServerInfo{
		Host:     rs,
		Port:     uint16(port),
		Param:    s.obfsParam,
		TcpMss:   1460,
		IVLen:    s.cipher.IVLen(),
		Key:      s.cipher.Key(),
		KeyLen:   s.cipher.KeyLen(),
		Overhead: overhead,
	})

	cipher, err := streamCipher.NewStreamCipher(obfs, s.encryptMethod, s.encryptPassword)
	if err != nil {
		return nil, err
	}

	protocol.SetServerInfo(&ssr.ServerInfo{
		Host:     rs,
		Port:     uint16(port),
		Param:    s.protocolParam,
		TcpMss:   1460,
		IV:       cipher.WriteIV(),
		IVLen:    len(cipher.WriteIV()),
		Key:      cipher.Key(),
		KeyLen:   len(cipher.Key()),
		Overhead: overhead,
	})

	if s.obfsData == nil {
		s.obfsData = obfs.GetData()
	}
	obfs.SetData(s.obfsData)

	if s.protocolData == nil {
		s.protocolData = protocol.GetData()
	}
	protocol.SetData(s.protocolData)

	// obfsServerInfo.SetHeadLen(b, 30)
	// protocolServerInfo.SetHeadLen(b, 30)

	ssrconn := NewSSTCPConn(cipher, protocol)
	if ssrconn.Conn == nil || ssrconn.RemoteAddr() == nil {
		return nil, errors.New("[ssr] nil connection")
	}

	target, err := socks5client.ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	if _, err := ssrconn.Write(target); err != nil {
		_ = ssrconn.Close()
		return nil, err
	}

	// log.Println("--------------return ssrconn --------------")
	return ssrconn, nil
}

func (s *Shadowsocksr) PacketConn(addr string) (net.PacketConn, error) {
	return net.ListenPacket("udp", "")
}
