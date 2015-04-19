package netfs

import (
	"os"
)


func (f *netFile) Chmod(mode os.FileMode) (err error) {
	defer onPanic(&err)

	f.doRequest(FILE_CHMOD)
	f.writeUint32(f.fid)
	f.writeUint32(uint32(mode))

	f.waitResponse(FILE_CHMOD)
	return f.readError()
}

func (f *Server) f_chmod() {
	fid := f.readUint32()
	mode := f.readUint32()

	err := f.getFile(fid).Chmod(os.FileMode(mode))

	f.doResponse(FILE_CHMOD)
	f.writeError(err)
}

//---------

func (f *netFile) Close() (err error) {
	defer onPanic(&err)

	f.doRequest(FILE_CLOSE)
	f.writeUint32(f.fid)

	f.waitResponse(FILE_CLOSE)
	return f.readError()
}

func (f *Server) f_close() {
	fid := f.readUint32()
	err := f.getFile(fid).Close()
	f.delFile(fid)

	f.doResponse(FILE_CLOSE)
	f.writeError(err)
}

//-----------

func (f *netFile) Read(b []byte) (n int, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_READ)
	f.writeUint32(f.fid)
	f.writeUint32(uint32(len(b)))

	f.waitResponse(FILE_READ)
	n = f.readByteTo(b)
	err = f.readError()
	return
}

func (f *Server) f_read() {
	fid := f.readUint32()
	_len := f.readUint32()

	b := make([]byte, int(_len))
	n, err := f.getFile(fid).Read(b)

	f.doResponse(FILE_READ)
	f.writeByte(b[:n])
	f.writeError(err)
}

//----------------

func (f *netFile) ReadAt(b []byte, off int64) (n int, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_READAT)
	f.writeUint32(f.fid)
	f.writeUint32(uint32(len(b)))
	f.writeInt64(off)

	f.waitResponse(FILE_READAT)
	n = f.readByteTo(b)
	err = f.readError()
	return
}

func (f *Server) f_readAt() {
	fid := f.readUint32()
	_len := f.readUint32()
	off := f.readInt64()

	b := make([]byte, int(_len))
	n, err := f.getFile(fid).ReadAt(b, off)

	f.doResponse(FILE_READAT)
	f.writeByte(b[:n])
	f.writeError(err)
}

//------------------

func (f *netFile) Readdir(n int) (fi []os.FileInfo, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_READDIR)
	f.writeUint32(f.fid)
	f.writeInt32(int32(n))

	f.waitResponse(FILE_READDIR)
	nu := f.readUint32()
	for i:=uint32(0); i<nu; i++ {
		_fi := f.readFileInfo()
		fi = append(fi, _fi)
	}

	err = f.readError()
	return
}

func (f *Server) f_readdir() {
	fid := f.readUint32()
	_n := f.readInt32()

	fis, err := f.getFile(fid).Readdir(int(_n))

	f.doResponse(FILE_READDIR)
	f.writeUint32(uint32(len(fis)))
	for _, fi := range fis {
		f.writeFileInfo(fi)
	}
	f.writeError(err)
}

//----------------

func (f *netFile) Readdirnames(n int) (names []string, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_READDIRNAMES)
	f.writeUint32(f.fid)
	f.writeInt32(int32(n))

	f.waitResponse(FILE_READDIRNAMES)
	nu := f.readUint32()
	for i:=uint32(0); i<nu; i++ {
		name := f.readString()
		names = append(names, name)
	}

	err = f.readError()
	return
}

func (f *Server) f_readdirnames() {
	fid := f.readUint32()
	_n := f.readInt32()

	names, err := f.getFile(fid).Readdirnames(int(_n))

	f.doResponse(FILE_READDIRNAMES)
	f.writeUint32(uint32(len(names)))
	for _, name := range names {
		f.writeString(name)
	}
	f.writeError(err)
}

//------------------

func (f *netFile) Seek(offset int64, whence int) (ret int64, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_SEEK)
	f.writeUint32(f.fid)
	f.writeInt64(offset)
	f.writeInt16(int16(whence))

	f.waitResponse(FILE_SEEK)
	ret = f.readInt64()
	err = f.readError()
	return
}

func (f *Server) f_seek() {
	fid := f.readUint32()
	offset := f.readInt64()
	_whence := f.readInt16()

	ret, err := f.getFile(fid).Seek(offset, int(_whence))

	f.doResponse(FILE_SEEK)
	f.writeInt64(ret)
	f.writeError(err)
}

//----------------

func (f *netFile) Stat() (fi os.FileInfo, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_STAT)
	f.writeUint32(f.fid)

	f.waitResponse(FILE_STAT)
	fi = f.readFileInfo()
	err = f.readError()
	return
}

func (f *Server) f_stat() {
	fid := f.readUint32()

	fi, err := f.getFile(fid).Stat()

	f.doResponse(FILE_STAT)
	f.writeFileInfo(fi)
	f.writeError(err)
}

//---------------

func (f *netFile) Sync() (err error) {
	defer onPanic(&err)

	f.doRequest(FILE_SYNC)
	f.writeUint32(f.fid)

	f.waitResponse(FILE_SYNC)
	return f.readError()
}

func (f *Server) f_sync() {
	fid := f.readUint32()

	err := f.getFile(fid).Sync()

	f.doResponse(FILE_SYNC)
	f.writeError(err)
}

//----------------

func (f *netFile) Truncate(size int64) (err error) {
	defer onPanic(&err)

	f.doRequest(FILE_TRUNCATE)
	f.writeUint32(f.fid)
	f.writeInt64(size)

	f.waitResponse(FILE_TRUNCATE)
	return f.readError()
}

func (f *Server) f_truncate() {
	fid := f.readUint32()
	size := f.readInt64()

	err := f.getFile(fid).Truncate(size)

	f.doResponse(FILE_TRUNCATE)
	f.writeError(err)
}

//-----------------

func (f *netFile) Write(b []byte) (n int, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_WRITE)
	f.writeUint32(f.fid)
	f.writeByte(b)

	f.waitResponse(FILE_WRITE)
	n = int(f.readUint32())
	err = f.readError()
	return
}

func (f *Server) f_write() {
	fid := f.readUint32()
	b := f.readByte()

	n, err := f.getFile(fid).Write(b)

	f.doResponse(FILE_WRITE)
	f.writeUint32(uint32(n))
	f.writeError(err)
}

//----------------

func (f *netFile) WriteAt(b []byte, off int64) (n int, err error) {
	defer onPanic(&err)

	f.doRequest(FILE_WRITEAT)
	f.writeUint32(f.fid)
	f.writeByte(b)
	f.writeInt64(off)

	f.waitResponse(FILE_WRITEAT)
	n = int(f.readUint32())
	err = f.readError()
	return
}

func (f *Server) f_writeAt() {
	fid := f.readUint32()
	b := f.readByte()
	off := f.readInt64()

	n, err := f.getFile(fid).WriteAt(b, off)

	f.doResponse(FILE_WRITEAT)
	f.writeUint32(uint32(n))
	f.writeError(err)
}


