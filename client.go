package stream

import (
	"os"
	"net"
	"time"
)

func Dial(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	c := new(Client)
	c.init(conn, 4096, 4096)
	return c, nil
}

type Client struct {
	conn
}

func (c *Client) Chmod(name string, mode os.FileMode) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_chmod(name, mode)
}

func (c *Client) Chtimes(name string, atime time.Time, mtime time.Time) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_chtimes(name, atime, mtime)
}

func (c *Client) Mkdir(name string, perm os.FileMode) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_mkdir(name, perm)
}

func (c *Client) MkdirAll(path string, perm os.FileMode) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_mkdirAll(path, perm)
}

func (c *Client) Remove(name string) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_remove(name)
}

func (c *Client) RemoveAll(path string) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_removeAll(path)
}

func (c *Client) Rename(oldpath, newpath string) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_rename(oldpath, newpath)
}

func (c *Client) Truncate(name string, size int64) (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_truncate(name, size)
}

func (c *Client) Create(name string) (file File, err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_create(name)
}

func (c *Client) Open(name string) (file File, err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_open(name)
}

func (c *Client) OpenFile(name string, flag int, perm os.FileMode) (file File, err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_openFile(name, flag, perm)
}

func (c *Client) Stat(name string) (fi os.FileInfo, err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_stat(name)
}

func (c *Client) Lstat(name string) (fi os.FileInfo, err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
	return c.fs_lstat(name)
}

func (c *Client) Close () (err error) {
	defer onPanic(&err)
	c.setDeadline(ActionTimeout)
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

func (f *netFile) Chmod(mode os.FileMode) (err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_chmod(mode)
}

func (f *netFile) Close() (err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_close()
}

func (f *netFile) Name() string {
	return f.name
}

func (f *netFile) Read(b []byte) (n int, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_read(b)
}

func (f *netFile) ReadAt(b []byte, off int64) (n int, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_readAt(b, off)
}

func (f *netFile) Readdir(n int) (fi []os.FileInfo, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_readdir(n)
}

func (f *netFile) Readdirnames(n int) (names []string, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_readdirnames(n)
}

func (f *netFile) Seek(offset int64, whence int) (ret int64, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_seek(offset, whence)
}

func (f *netFile) Stat() (fi os.FileInfo, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_stat()
}

func (f *netFile) Sync() (err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_sync()
}

func (f *netFile) Truncate(size int64) (err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_truncate(size)
}

func (f *netFile) Write(b []byte) (n int, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_write(b)
}

func (f *netFile) WriteAt(b []byte, off int64) (n int, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_writeAt(b, off)
}

func (f *netFile) WriteString(s string) (ret int, err error) {
	defer onPanic(&err)
	f.conn.setDeadline(ActionTimeout)
	return f.f_write([]byte(s))
}

func onPanic(err *error) {
	if x := recover(); x != nil {
		switch v := x.(type) {
		case IO_Error :
			*err = v
		default:
			panic(x)
		}
	}
}

