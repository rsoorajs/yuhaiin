package obfs

import (
	"math/rand"
	"net"
)

func init() {
	register("http_post", newHttpPost)
}

// newHttpPost create a http_post object
func newHttpPost(con net.Conn) IObfs {
	// newHttpSimple create a http_simple object

	t := &httpSimplePost{
		userAgentIndex: rand.Intn(len(requestUserAgent)),
		methodGet:      false,
		Conn:           con,
	}
	return t
}
