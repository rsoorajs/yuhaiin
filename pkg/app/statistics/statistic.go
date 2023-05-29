package statistics

import (
	"context"
	"fmt"
	"net"

	"github.com/Asutorufa/yuhaiin/pkg/app/shunt"
	"github.com/Asutorufa/yuhaiin/pkg/log"
	proxy "github.com/Asutorufa/yuhaiin/pkg/net/interfaces"
	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/direct"
	"github.com/Asutorufa/yuhaiin/pkg/protos/config/listener"
	"github.com/Asutorufa/yuhaiin/pkg/protos/statistic"
	gs "github.com/Asutorufa/yuhaiin/pkg/protos/statistic/grpc"
	"github.com/Asutorufa/yuhaiin/pkg/utils/cache"
	"github.com/Asutorufa/yuhaiin/pkg/utils/convert"
	"github.com/Asutorufa/yuhaiin/pkg/utils/goos"
	"github.com/Asutorufa/yuhaiin/pkg/utils/id"
	"github.com/Asutorufa/yuhaiin/pkg/utils/slice"
	"github.com/Asutorufa/yuhaiin/pkg/utils/syncmap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Connections struct {
	gs.UnimplementedConnectionsServer

	proxy.Proxy
	idSeed id.IDGenerator

	connStore syncmap.SyncMap[uint64, connection]

	processDumper listener.ProcessDumper
	Cache         *Cache

	notify notify
}

func NewConnStore(cache *cache.Cache, dialer proxy.Proxy, processDumper listener.ProcessDumper) *Connections {
	if dialer == nil {
		dialer = direct.Default
	}

	c := &Connections{
		Proxy:         dialer,
		processDumper: processDumper,
		Cache:         NewCache(cache),
	}

	return c
}

func (c *Connections) Notify(_ *emptypb.Empty, s gs.Connections_NotifyServer) error {
	id := c.notify.register(s, c.connStore.ValueSlice()...)
	defer c.notify.unregister(id)
	log.Debug("new notify client", "id", id)
	<-s.Context().Done()
	log.Debug("remove notify client", "id", id)
	return s.Context().Err()
}

func (c *Connections) Conns(context.Context, *emptypb.Empty) (*gs.ConnectionsInfo, error) {
	return &gs.ConnectionsInfo{
		Connections: slice.To(c.connStore.ValueSlice(),
			func(c connection) *statistic.Connection { return c.Info() }),
	}, nil
}

func (c *Connections) CloseConn(_ context.Context, x *gs.ConnectionsId) (*emptypb.Empty, error) {
	for _, x := range x.Ids {
		if z, ok := c.connStore.Load(x); ok {
			z.Close()
		}
	}
	return &emptypb.Empty{}, nil
}

func (c *Connections) Close() error {
	c.connStore.Range(func(key uint64, v connection) bool {
		v.Close()
		return true
	})

	c.Cache.Close()
	return nil
}

func (c *Connections) Total(context.Context, *emptypb.Empty) (*gs.TotalFlow, error) {
	return &gs.TotalFlow{
		Download: c.Cache.LoadDownload(),
		Upload:   c.Cache.LoadUpload(),
	}, nil
}

func (c *Connections) Remove(id uint64) {
	if z, ok := c.connStore.LoadAndDelete(id); ok {
		log.Debug("close conn",
			"id", z.ID(),
			"addr", z.Info().Addr,
			"s0urce", z.Info().Extra[(proxy.SourceKey{}).String()],
			"outbound", getRemote(z))
	}

	c.notify.pubRemoveConns(id)
}

func (c *Connections) storeConnection(o connection) {
	c.connStore.Store(o.ID(), o)
	c.notify.pubNewConns(o)
	log.Debug("new conn",
		"id", o.ID(),
		"addr", o.Info().Addr,
		"network", o.Info().Type.ConnType,
		"outbound", o.Info().Extra["Outbound"])
}

func (c *Connections) PacketConn(ctx context.Context, addr proxy.Address) (net.PacketConn, error) {
	process := c.DumpProcess(ctx, addr)
	con, err := c.Proxy.PacketConn(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("dial packet conn (%s) failed: %w", process, err)
	}

	z := &packetConn{con, c.getConnection(ctx, con, addr), c}

	c.storeConnection(z)
	return z, nil
}

func getRemote(con any) string {
	r, ok := con.(interface{ RemoteAddr() net.Addr })
	if ok {
		return r.RemoteAddr().String()
	}

	return ""
}

func getRealAddr(store proxy.Store, addr proxy.Address) string {
	z, ok := store.Get(shunt.DOMAIN_MARK_KEY{})
	if ok {
		s, ok := convert.ToString(z)
		if ok {
			return s
		}
	}

	return addr.String()
}

func (c *Connections) getConnection(ctx context.Context, conn interface{ LocalAddr() net.Addr }, addr proxy.Address) *statistic.Connection {
	store := proxy.StoreFromContext(ctx)

	connection := &statistic.Connection{
		Id:   c.idSeed.Generate(),
		Addr: getRealAddr(store, addr),
		Type: &statistic.NetType{
			ConnType:       addr.NetworkType(),
			UnderlyingType: statistic.Type(statistic.Type_value[conn.LocalAddr().Network()]),
		},
		Extra: convert.ToStringMap(store),
	}

	if out := getRemote(conn); out != "" {
		connection.Extra["Outbound"] = out
	}
	return connection
}

func (c *Connections) Conn(ctx context.Context, addr proxy.Address) (net.Conn, error) {
	process := c.DumpProcess(ctx, addr)
	con, err := c.Proxy.Conn(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("dial conn (%s) failed: %w", process, err)
	}

	z := &conn{con, c.getConnection(ctx, con, addr), c}
	c.storeConnection(z)
	return z, nil
}

func (c *Connections) DumpProcess(ctx context.Context, addr proxy.Address) (s string) {
	if c.processDumper == nil {
		return
	}

	store := proxy.StoreFromContext(ctx)

	source, ok := store.Get(proxy.SourceKey{})
	if !ok {
		return
	}

	var dst any
	if goos.IsAndroid == 1 {
		dst, ok = store.Get(proxy.InboundKey{})
		if !ok {
			dst, ok = store.Get(proxy.DestinationKey{})
		}
	} else {
		dst, ok = store.Get(proxy.DestinationKey{})
	}
	if !ok {
		return
	}

	sourceAddr, err := convert.ToProxyAddress(addr.NetworkType(), source)
	if err != nil {
		return
	}

	dstAddr, err := convert.ToProxyAddress(addr.NetworkType(), dst)
	if err != nil {
		return
	}

	process, err := c.processDumper.ProcessName(addr.Network(), sourceAddr, dstAddr)
	if err != nil {
		log.Warn("dump process failed", "err", err)
		return
	}

	store.Add("Process", process)
	return process
}
