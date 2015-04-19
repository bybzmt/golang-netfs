package stream

import (
	"os"
	"time"
)

func (c *Client) fs_chmod(name string, mode os.FileMode) error {
	c.doRequest(FS_CHMOD)
	c.writeString(name)
	c.writeUint32(uint32(mode))

	c.waitResponse(FS_CHMOD)
	return c.readError()
}

func (c *Server) fs_chmod() {
	name := c.readString()
	mode := c.readUint32()

	err := c.fs.Chmod(name, os.FileMode(mode))

	c.doResponse(FS_CHMOD)
	c.writeError(err)
}

//----------------------

func (c *Client) fs_chtimes(name string, atime time.Time, mtime time.Time) error {
	c.doRequest(FS_CHTIMES)
	c.writeString(name)
	c.writeInt64(atime.Unix())
	c.writeInt64(mtime.Unix())

	c.waitResponse(FS_CHTIMES)
	return c.readError()
}

func (c *Server) fs_chtimes() {
	name := c.readString()
	t1 := c.readInt64()
	t2 := c.readInt64()

	err := c.fs.Chtimes(name, time.Unix(t1, 0), time.Unix(t2, 0))

	c.doResponse(FS_CHTIMES)
	c.writeError(err)
}

//------------------

func (c *Client) fs_mkdir(name string, perm os.FileMode) error {
	c.doRequest(FS_MKDIR)
	c.writeString(name)
	c.writeUint32(uint32(perm))

	c.waitResponse(FS_MKDIR)
	return c.readError()
}

func (c *Server) fs_mkdir() {
	name := c.readString()
	perm := c.readUint32()

	err := c.fs.Mkdir(name, os.FileMode(perm))

	c.doResponse(FS_MKDIR)
	c.writeError(err)
}

//---------------

func (c *Client) fs_mkdirAll(path string, perm os.FileMode) error {
	c.doRequest(FS_MKDIRALL)
	c.writeString(path)
	c.writeUint32(uint32(perm))

	c.waitResponse(FS_MKDIRALL)
	return c.readError()
}

func (c *Server) fs_mkdirAll() {
	path := c.readString()
	perm := c.readUint32()

	err := c.fs.MkdirAll(path, os.FileMode(perm))

	c.doResponse(FS_MKDIRALL)
	c.writeError(err)
}

//----------------

func (c *Client) fs_remove(name string) error {
	c.doRequest(FS_REMOVE)
	c.writeString(name)

	c.waitResponse(FS_REMOVE)
	return c.readError()
}

func (c *Server) fs_remove() {
	name := c.readString()

	err := c.fs.Remove(name)

	c.doResponse(FS_REMOVE)
	c.writeError(err)
}

//-------------

func (c *Client) fs_removeAll(path string) error {
	c.doRequest(FS_REMOVEALL)
	c.writeString(path)

	c.waitResponse(FS_REMOVEALL)
	return c.readError()
}

func (c *Server) fs_removeAll() {
	path := c.readString()

	err := c.fs.RemoveAll(path)

	c.doResponse(FS_REMOVEALL)
	c.writeError(err)
}

//-------------

func (c *Client) fs_rename(oldpath, newpath string) error {
	c.doRequest(FS_RENAME)
	c.writeString(oldpath)
	c.writeString(newpath)

	c.waitResponse(FS_RENAME)
	return c.readError()
}

func (c *Server) fs_rename() {
	oldpath := c.readString()
	newpath := c.readString()

	err := c.fs.Rename(oldpath, newpath)

	c.doResponse(FS_RENAME)
	c.writeError(err)
}

//-------------

func (c *Client) fs_truncate(name string, size int64) error {
	c.doRequest(FS_TRUNCATE)
	c.writeString(name)
	c.writeInt64(size)

	c.waitResponse(FS_TRUNCATE)
	return c.readError()
}

func (c *Server) fs_truncate() {
	name := c.readString()
	size := c.readInt64()

	err := c.fs.Truncate(name, size)

	c.doResponse(FS_TRUNCATE)
	c.writeError(err)
}

//----------------

func _client_init_file(c *Client, fid uint32, name string, err error) (File, error) {
	if err != nil {
		return nil, err
	}

	f := new(netFile)
	f.Client = c
	f.fid = fid
	f.name = name

	return f, nil
}

//-----------------

func (c *Client) fs_create(name string) (file File, err error) {
	c.doRequest(FS_CREATE)
	c.writeString(name)

	c.waitResponse(FS_CREATE)
	fid := c.readUint32()
	_err := c.readError()

	return _client_init_file(c, fid, name, _err)
}

func (c *Server) fs_create() {
	name := c.readString()

	fd, err := c.fs.Create(name)
	fid := c.addFile(fd)

	c.doResponse(FS_CREATE)
	c.writeUint32(fid)
	c.writeError(err)
}
//-----------------

func (c *Client) fs_open(name string) (file File, err error) {
	c.doRequest(FS_OPEN)
	c.writeString(name)

	c.waitResponse(FS_OPEN)
	fid := c.readUint32()
	_err := c.readError()

	return _client_init_file(c, fid, name, _err)
}

func (c *Server) fs_open() {
	name := c.readString()

	fd, err := c.fs.Open(name)
	fid := c.addFile(fd)

	c.doResponse(FS_OPEN)
	c.writeUint32(fid)
	c.writeError(err)
}

//-----------------

func (c *Client) fs_openFile(name string, flag int, perm os.FileMode) (file File, err error) {
	c.doRequest(FS_OPENFILE)
	c.writeString(name)
	c.writeInt32(int32(flag))
	c.writeUint32(uint32(perm))

	c.waitResponse(FS_OPENFILE)
	fid := c.readUint32()
	_err := c.readError()

	return _client_init_file(c, fid, name, _err)
}

func (c *Server) fs_openFile() {
	name := c.readString()
	flag := c.readInt32()
	perm := c.readUint32()

	fd, err := c.fs.OpenFile(name, int(flag), os.FileMode(perm))
	fid := c.addFile(fd)

	c.doResponse(FS_OPENFILE)
	c.writeUint32(fid)
	c.writeError(err)
}

//------

func (c *Client) fs_lstat(name string) (fi os.FileInfo, err error) {
	c.doRequest(FS_LSTAT)
	c.writeString(name)

	c.waitResponse(FS_LSTAT)
	fi = c.readFileInfo()
	err = c.readError()
	return
}

func (c *Server) fs_lstat() {
	name := c.readString()

	fi, err := c.fs.Lstat(name)

	c.doResponse(FS_LSTAT)
	c.writeFileInfo(fi)
	c.writeError(err)
}

//------

func (c *Client) fs_stat(name string) (fi os.FileInfo, err error) {
	c.doRequest(FS_STAT)
	c.writeString(name)

	c.waitResponse(FS_STAT)
	fi = c.readFileInfo()
	err = c.readError()
	return
}

func (c *Server) fs_stat() {
	name := c.readString()

	fi, err := c.fs.Stat(name)

	c.doResponse(FS_STAT)
	c.writeFileInfo(fi)
	c.writeError(err)
}

