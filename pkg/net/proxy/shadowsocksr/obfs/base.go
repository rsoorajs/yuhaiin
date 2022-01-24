package obfs

import (
	"net"
	"strings"

	ssr "github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/utils"
)

type creator func(net.Conn) IObfs

var (
	creatorMap = make(map[string]creator)
)

type IObfs interface {
	SetServerInfo(s *ssr.ServerInfo)
	GetServerInfo() (s *ssr.ServerInfo)
	Encode(data []byte) (encodedData []byte, err error)
	Decode(data []byte) (decodedData []byte, needSendBack bool, err error)
	SetData(data interface{})
	GetData() interface{}
	GetOverhead() int

	net.Conn
}

func register(name string, c creator) {
	creatorMap[name] = c
}

// NewObfs create an obfs object by name and return as an IObfs interface
func NewObfs(conn net.Conn, name string) IObfs {
	c, ok := creatorMap[strings.ToLower(name)]
	if ok {
		return c(conn)
	}
	return nil
}
