package netfs

import (
	"testing"
	"os"
	"bytes"
)

type testFile struct {
	t *testing.T
	mode os.FileMode
	expect_call string
	call string
	err error
	n int
	whence int
	ret int64
	off int64
	offset int64
	size int64
	b []byte
	fi os.FileInfo
	fis []os.FileInfo
	dirs []string
}

func (f *testFile) Chmod(mode os.FileMode) (err error) {
	f.call = "chmod"

	if mode != f.mode {
		f.t.Errorf("Chmod expect:%s give:%s", f.mode, mode)
	}

	return f.err
}

func (f *testFile) Close() (err error) {
	f.call = "close"

	return f.err
}

func (f *testFile) Name() string {
	return ""
}

func (f *testFile) Read(b []byte) (n int, err error) {
	f.call = "read"

	r := bytes.NewReader(f.b)

	return r.Read(b)
}

func (f *testFile) ReadAt(b []byte, off int64) (n int, err error) {
	f.call = "readat"

	r := bytes.NewReader(f.b)

	return r.ReadAt(b, off)
}

func (f *testFile) Readdir(n int) (fi []os.FileInfo, err error) {
	f.call = "readat"

	if n != f.n {
		f.t.Errorf("Readdirnames expect:%d give:%s", f.n, n)
	}

	return f.fis, f.err
}

func (f *testFile) Readdirnames(n int) (names []string, err error) {
	f.call = "readdirnames"

	if n != f.n {
		f.t.Errorf("Readdirnames expect:%d give:%s", f.n, n)
	}

	return f.dirs, f.err
}

func (f *testFile) Seek(offset int64, whence int) (ret int64, err error) {
	f.call = "seek"

	if offset != f.offset {
		f.t.Errorf("Seek expect:%d give:%s", f.offset, offset)
	}
	if whence != f.whence {
		f.t.Errorf("Seek expect:%d give:%s", f.whence, whence)
	}

	return f.ret, f.err
}

func (f *testFile) Stat() (fi os.FileInfo, err error) {
	f.call = "stat"

	return f.fi, f.err
}

func (f *testFile) Sync() (err error) {
	f.call = "sync"

	return f.err
}

func (f *testFile) Truncate(size int64) (err error) {
	f.call = "truncate"

	if size != f.size {
		f.t.Errorf("Truncate expect:%d give:%s", f.size, size)
	}

	return f.err
}

func (f *testFile) Write(b []byte) (n int, err error) {
	f.call = "write"

	if string(b) != string(f.b) {
		f.t.Errorf("Write expect:%s give:%s", string(f.b), string(b))
	}

	return f.n, f.err
}

func (f *testFile) WriteAt(b []byte, off int64) (n int, err error) {
	f.call = "writeat"

	if string(b) != string(f.b) {
		f.t.Errorf("WriteAt expect:%s give:%s", string(f.b), string(b))
	}
	if off != f.off {
		f.t.Errorf("WriteAt expect:%d give:%s", f.off, off)
	}

	return f.n, f.err
}

func (f *testFile) WriteString(s string) (ret int, err error) {
	return
}
