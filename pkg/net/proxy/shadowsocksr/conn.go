package shadowsocksr

import (
	"bytes"
	"net"

	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocksr/protocol"
	"github.com/Asutorufa/yuhaiin/pkg/net/utils"
)

// SSTCPConn the struct that override the net.Conn methods
type SSTCPConn struct {
	net.Conn
	IProtocol           protocol.IProtocol
	readBuf             []byte
	underPostdecryptBuf *bytes.Buffer
	decryptedBuf        *bytes.Buffer
}

func NewSSTCPConn(c net.Conn, protocol protocol.IProtocol) *SSTCPConn {
	return &SSTCPConn{
		Conn:                c,
		IProtocol:           protocol,
		readBuf:             *utils.BuffPool(utils.DefaultSize).Get().(*[]byte),
		decryptedBuf:        new(bytes.Buffer),
		underPostdecryptBuf: new(bytes.Buffer),
	}
}

func (c *SSTCPConn) Close() error {
	utils.BuffPool(utils.DefaultSize).Put(&c.readBuf)
	return c.Conn.Close()
}

// func (c *SSTCPConn) initEncryptor(b []byte) (iv []byte, err error) {
// 	if !c.EncryptInited {
// 		var key []byte
// 		if c, ok := c.Conn.(*streamCipher.StreamCipher); ok {
// 			iv = c.WriteIV()
// 			key = c.Key()
// 		} else {
// 			return nil, fmt.Errorf("cipher not streamCipher")
// 		}
// 		// iv, err = c.WriteIV()
// 		// if err != nil {
// 		// return nil, err
// 		// }
// 		overhead := c.IObfs.GetOverhead() + c.IProtocol.GetOverhead()
// 		// should initialize obfs/protocol now, because iv is ready now
// 		obfsServerInfo := c.IObfs.GetServerInfo()
// 		obfsServerInfo.SetHeadLen(b, 30)
// 		obfsServerInfo.IV, obfsServerInfo.IVLen = iv, len(iv)
// 		obfsServerInfo.Key, obfsServerInfo.KeyLen = key, len(key)
// 		obfsServerInfo.Overhead = overhead
// 		c.IObfs.SetServerInfo(obfsServerInfo)

// 		protocolServerInfo := c.IProtocol.GetServerInfo()
// 		protocolServerInfo.SetHeadLen(b, 30)
// 		protocolServerInfo.IV, protocolServerInfo.IVLen = iv, len(iv)
// 		protocolServerInfo.Key, protocolServerInfo.KeyLen = key, len(key)
// 		protocolServerInfo.Overhead = overhead
// 		c.IProtocol.SetServerInfo(protocolServerInfo)
// 	}
// 	return
// }

func (c *SSTCPConn) Read(b []byte) (n int, err error) {
	for {
		n, err = c.doRead(b)
		if b == nil || n != 0 || err != nil {
			return n, err
		}
	}
}

func (c *SSTCPConn) doRead(b []byte) (n int, err error) {
	//先吐出已经解密后数据
	if c.decryptedBuf.Len() > 0 {
		return c.decryptedBuf.Read(b)
	}

	n, err = c.Conn.Read(c.readBuf)
	if n == 0 || err != nil {
		return n, err
	}

	c.underPostdecryptBuf.Write(c.readBuf[:n])
	// and read it to buf immediately
	buf := c.underPostdecryptBuf.Bytes()
	postDecryptedData, length, err := c.IProtocol.PostDecrypt(buf)
	if err != nil {
		c.underPostdecryptBuf.Reset()
		//log.Println(string(decodebytes))
		//log.Println("err", err)
		return 0, err
	}

	if length == 0 {
		// not enough to postDecrypt
		return 0, nil
	}

	c.underPostdecryptBuf.Next(length)

	postDecryptedLength := len(postDecryptedData)
	blength := len(b)
	//b的长度是否够用
	if blength >= postDecryptedLength {
		copy(b, postDecryptedData)
		return postDecryptedLength, nil
	}
	copy(b, postDecryptedData[:blength])
	c.decryptedBuf.Write(postDecryptedData[blength:])
	return blength, nil
}

func (c *SSTCPConn) Write(b []byte) (n int, err error) {
	outData, err := c.IProtocol.PreEncrypt(b)
	if err != nil {
		return 0, err
	}
	_, err = c.Conn.Write(outData)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
