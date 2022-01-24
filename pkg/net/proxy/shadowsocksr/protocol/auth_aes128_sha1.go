package protocol

import (
	"bytes"

	ssr "github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/utils"
)

func init() {
	register("auth_aes128_sha1", NewAuthAES128SHA1)
}

func NewAuthAES128SHA1() IProtocol {
	a := &authAES128{
		salt:       "auth_aes128_sha1",
		hmac:       ssr.HmacSHA1,
		hashDigest: ssr.SHA1Sum,
		packID:     1,
		recvInfo: recvInfo{
			recvID: 1,
			buffer: bytes.NewBuffer(nil),
		},
	}
	return a
}
