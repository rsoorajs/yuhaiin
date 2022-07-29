package statistic

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	ylog "github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/interfaces/proxy"
	"github.com/Asutorufa/yuhaiin/pkg/protos/config"
	grpcsts "github.com/Asutorufa/yuhaiin/pkg/protos/grpc/statistic"
	"github.com/Asutorufa/yuhaiin/pkg/protos/statistic"
	"github.com/Asutorufa/yuhaiin/pkg/utils/syncmap"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type conns interface {
	AddConn(_ net.Conn, host proxy.Address) net.Conn
	AddPacketConn(_ net.PacketConn, host proxy.Address) net.PacketConn
}

var _ conns = (*counter)(nil)

type counter struct {
	grpcsts.UnimplementedConnectionsServer

	accountant

	idSeed idGenerater
	conns  syncmap.SyncMap[int64, connection]
}

func newStatistic() *counter { return &counter{} }

func (c *counter) Conns(context.Context, *emptypb.Empty) (*statistic.ConnResp, error) {
	resp := &statistic.ConnResp{}
	c.conns.Range(func(key int64, v connection) bool {
		resp.Connections = append(resp.Connections, v.Info())
		return true
	})

	return resp, nil
}

func (c *counter) CloseConn(_ context.Context, x *statistic.CloseConnsReq) (*emptypb.Empty, error) {
	for _, x := range x.Conns {
		if z, ok := c.conns.Load(x); ok {
			z.Close()
		}
	}
	return &emptypb.Empty{}, nil
}

func (c *counter) CloseAll() {
	c.conns.Range(func(key int64, v connection) bool {
		v.Close()
		return true
	})
}

func (c *counter) Statistic(_ *emptypb.Empty, srv grpcsts.Connections_StatisticServer) error {
	log.Println("Start Send Flow Message to Client.")
	id := c.accountant.AddClient(srv.Send)
	<-srv.Context().Done()
	c.accountant.RemoveClient(id)
	log.Println("Client is Hidden, Close Stream.")
	return srv.Context().Err()
}

func (c *counter) delete(id int64) {
	if z, ok := c.conns.LoadAndDelete(id); ok {
		ylog.Debugln("close", c.cString(z))
	}
}

func (c *counter) storeConnection(o connection) {
	ylog.Debugf(c.cString(o))
	c.conns.Store(o.GetId(), o)
}

func (c *counter) cString(o connection) (s string) {
	if ylog.IsOutput(config.Logcat_debug) {
		s = fmt.Sprintf("%v| <%s[%v]>: %v(%s), %s <-> %s",
			o.GetId(), o.GetType(), o.GetExtra()[MODE_MARK], o.GetAddr(), getExtra(o), o.GetLocal(), o.GetRemote())
	}
	return
}

func getExtra(o connection) string {
	str := strings.Builder{}

	for k, v := range o.GetExtra() {
		if k == MODE_MARK {
			continue
		}
		str.WriteString(fmt.Sprintf("%s: %s,", k, v))
	}

	return str.String()
}
