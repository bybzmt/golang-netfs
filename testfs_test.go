package stream

import (
	"testing"
	"os"
	"time"
	"log"
)

type testFs struct {
	t *testing.T
	expect_call string
	call string
	name string
	oldpath string
	newpath string
	mode os.FileMode
	perm os.FileMode
	err error
	file File
	atime time.Time
	mtime time.Time
	size int64
	flag int
	fi *FileInfo
}

func (l *testFs) _test() {
	log.Println("ss")
}

func (l *testFs) Chmod(name string, mode os.FileMode) error {
	l.call = "chmod"

	if name != l.name {
		l.t.Errorf("Chomd name expect:%s give:%s", l.name, name)
	}
	if mode != l.mode {
		l.t.Errorf("Chomd mode expect:%o give:%o", l.mode, mode)
	}

	return l.err
}

func (l *testFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	l.call = "chtimes"

	if name != l.name {
		l.t.Errorf("Chtimes name expect:%s give:%s", l.name, name)
	}

	if !atime.Equal(l.atime) {
		l.t.Errorf("Chtimes atime expect:%s give:%s", l.atime.Format(time.RFC3339Nano), atime.Format(time.RFC3339Nano))
	}

	if !mtime.Equal(l.mtime) {
		l.t.Errorf("Chtimes mtime expect:%s give:%s", l.mtime.Format(time.RFC3339Nano), mtime.Format(time.RFC3339Nano))
	}

	return l.err
}

func (l *testFs) Mkdir(name string, perm os.FileMode) error {
	l.call = "mkdir"

	if name != l.name {
		l.t.Errorf("Mkdir name expect:%s give:%s", l.name, name)
	}
	if perm != l.perm {
		l.t.Errorf("Mkdir perm expect:%o give:%o", l.perm, perm)
	}

	return l.err
}

func (l *testFs) MkdirAll(name string, perm os.FileMode) error {
	l.call = "mkdirall"

	if name != l.name {
		l.t.Errorf("MkdirAll name expect:%s give:%s", l.name, name)
	}
	if perm != l.perm {
		l.t.Errorf("MkdirAll perm expect:%o give:%o", l.perm, perm)
	}

	return l.err
}

func (l *testFs) Remove(name string) error {
	l.call = "remove"

	if name != l.name {
		l.t.Errorf("Remove name expect:%s give:%s", l.name, name)
	}

	return l.err
}

func (l *testFs) RemoveAll(name string) error {
	l.call = "removeall"

	if name != l.name {
		l.t.Errorf("RemoveAll name expect:%s give:%s", l.name, name)
	}

	return l.err
}

func (l *testFs) Rename(oldpath, newpath string) error {
	l.call = "rename"

	if oldpath != l.oldpath {
		l.t.Errorf("Rename oldpath expect:%s give:%s", l.oldpath, oldpath)
	}
	if newpath != l.newpath {
		l.t.Errorf("Rename newpath expect:%s give:%s", l.newpath, newpath)
	}

	return l.err
}

func (l *testFs) Truncate(name string, size int64) error {
	l.call = "truncate"

	if name != l.name {
		l.t.Errorf("Trucate name expect:%s give:%s", l.name, name)
	}
	if size != l.size {
		l.t.Errorf("Trucate size expect:%d give:%d", l.size, size)
	}

	return l.err
}

func (l *testFs) Create(name string) (file File, err error) {
	l.call = "create"

	if name != l.name {
		l.t.Errorf("Create name expect:%s give:%s", l.name, name)
	}

	return l.file, l.err
}

func (l *testFs) Open(name string) (file File, err error) {
	l.call = "open"

	if name != l.name {
		l.t.Errorf("Open name expect:%s give:%s", l.name, name)
	}

	return l.file, l.err
}

func (l *testFs) OpenFile(name string, flag int, perm os.FileMode) (file File, err error) {
	l.call = "openfile"

	if name != l.name {
		l.t.Errorf("OpenFile name expect:%s give:%s", l.name, name)
	}
	if flag != l.flag {
		l.t.Errorf("OpenFile flag expect:%d give:%d", l.flag, flag)
	}
	if perm != l.perm {
		l.t.Errorf("OpenFile perm expect:%o give:%o", l.perm, perm)
	}

	return l.file, l.err
}

func (l *testFs) Stat(name string) (fi os.FileInfo, err error) {
	l.call = "stat"

	if name != l.name {
		l.t.Errorf("Stat name expect:%s give:%s", l.name, name)
	}

	return l.fi, l.err
}

func (l *testFs) Lstat(name string) (fi os.FileInfo, err error) {
	l.call = "lstat"

	if name != l.name {
		l.t.Errorf("Lstat name expect:%s give:%s", l.name, name)
	}

	return l.fi, l.err
}

