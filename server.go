package netfs

import (
	"net"
	"time"
	"fmt"
)

func Listen(addr string, fs FileSystem) {
	if fs == nil {
		fs = new(LocalFs).Init("./")
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		getLog().Crit(err.Error())
	}

	l := ln.(*net.TCPListener)

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			getLog().Crit(err.Error())
		}

		go RunRev(conn, fs)
	}
}

func RunRev(conn *net.TCPConn, fs FileSystem) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case IO_Error :
				getLog().Notice("io: " + conn.RemoteAddr().String() + " " + v.Error())
			case Data_Error :
				getLog().Notice("data: " + conn.RemoteAddr().String() + " " + v.Error())
			default:
				panic(v)
			}
		}
	}()

	getLog().Debug("link: from " + conn.RemoteAddr().String())
	defer getLog().Debug("link: close " + conn.RemoteAddr().String())

	conn.SetLinger(5)
	conn.SetKeepAlivePeriod(10 * time.Second)
	conn.SetKeepAlive(true)

	c := new(Server)
	c.init(conn, 4096, 4096)
	c.fs = fs
	c.fds = make(map[uint32]File)
	c.Run()
}

type Server struct {
	conn
	fs FileSystem
	fid uint32
	fds map[uint32]File
}

func (c *Server) Run() {
	defer c.Close()

	if err := c.LinkInit(); err != nil {
		getLog().Info("link: " + err.Error())
		return
	}

	for {
		c.setDeadline(IdleTimeout)

		code := c.waitRequest()
		switch code {
		case LINK_CLOSE :
			return
		case LINK_PING :
			c.doResponse(LINK_PING)
		default:
			c.doAction(code)
		}

		c.flush()
	}
}

func (c *Server) doAction(code uint8) {
	switch (code) {
		//文件系统操作码
		case FS_CHMOD     : c.fs_chmod()
		case FS_CHTIMES   : c.fs_chtimes()
		case FS_MKDIR     : c.fs_mkdir()
		case FS_MKDIRALL  : c.fs_mkdirAll()
		case FS_REMOVE    : c.fs_remove()
		case FS_REMOVEALL : c.fs_removeAll()
		case FS_RENAME    : c.fs_rename()
		case FS_TRUNCATE   : c.fs_truncate()
		case FS_CREATE    : c.fs_create()
		case FS_OPEN      : c.fs_open()
		case FS_OPENFILE  : c.fs_openFile()
		case FS_LSTAT     : c.fs_lstat()
		case FS_STAT      : c.fs_stat()

		//文件对象操作码
		case FILE_CHMOD    : c.f_chmod()
		case FILE_CLOSE    : c.f_close()
		case FILE_READ     : c.f_read()
		case FILE_READAT   : c.f_readAt()
		case FILE_READDIR  : c.f_readdir()
		case FILE_READDIRNAMES : c.f_readdirnames()
		case FILE_SEEK     : c.f_seek()
		case FILE_STAT     : c.f_stat()
		case FILE_SYNC     : c.f_sync()
		case FILE_TRUNCATE : c.f_truncate()
		case FILE_WRITE    : c.f_write()
		case FILE_WRITEAT  : c.f_writeAt()
		default:
		panic(Data_Error(fmt.Sprintf("Unexpect Target:%d", code)))
	}
}

func (c *Server) addFile(f File) uint32 {
	if f == nil {
		return 0
	}

	c.fid++
	c.fds[c.fid] = f
	return c.fid
}

func (c *Server) getFile(fid uint32) File {
	f, ok := c.fds[fid]
	if !ok {
		panic(Data_Error(fmt.Sprintf("Undefined Fid:%d", fid)))
	}
	return f
}

func (c *Server) delFile(fid uint32) {
	delete(c.fds, fid)
}

func (c *Server) Close() {
	for _, f := range c.fds {
		f.Close()
	}

	c.conn.Close()
}


