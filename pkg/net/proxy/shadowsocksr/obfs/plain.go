package obfs

import (
	"net"

	ssr "github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/utils"
)

func init() {
	register("plain", newPlainObfs)
}

type plain struct {
	ssr.ServerInfo
	net.Conn
}

func newPlainObfs(conn net.Conn) IObfs {
	p := &plain{Conn: conn}
	return p
}

func (p *plain) SetServerInfo(s *ssr.ServerInfo) {
	p.ServerInfo = *s
}

func (p *plain) GetServerInfo() (s *ssr.ServerInfo) {
	return &p.ServerInfo
}

func (p *plain) Encode(data []byte) (encodedData []byte, err error) {
	return data, nil
}

func (p *plain) Decode(data []byte) (decodedData []byte, needSendBack bool, err error) {
	return data, false, nil
}

func (p *plain) SetData(data interface{}) {

}

func (p *plain) GetData() interface{} {
	return nil
}

func (p *plain) GetOverhead() int {
	return 0
}
