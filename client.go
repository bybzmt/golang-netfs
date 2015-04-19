package netfs

import (
	"net"
)

func Dial(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return new(Client).Init(conn, 4096, 4096)
}

type Client struct {
	conn
}

func (c *Client) Init(conn net.Conn, rBuf, wBuf int) (*Client, error) {
	c.init(conn, rBuf, wBuf)

	if err := c.LinkInit(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close () (err error) {
	defer onPanic(&err)
	c.doRequest(LINK_CLOSE)
	c.flush()
	c.conn.Close()
	return
}

type netFile struct {
	*Client
	name string
	fid uint32
}

func (f *netFile) Name() string {
	return f.name
}

func (f *netFile) WriteString(s string) (ret int, err error) {
	return f.Write([]byte(s))
}
