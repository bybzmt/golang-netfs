package netfs

import (
	"testing"
	"os"
	"io"
	"log"
	"errors"
	"time"
)

var test_client *Client
var test_client_file File
var test_fs *testFs
var test_file *testFile

func TestMain(m *testing.M) {
	test_fs = new(testFs)
	go Listen("127.0.0.1:11120", test_fs)

	time.Sleep(300 * time.Millisecond)

	var err error
	test_client, err = Dial("127.0.0.1:11120")
	if err != nil {
		log.Fatal(err)
	}

	ex := m.Run()

	time.Sleep(300 * time.Millisecond)

	os.Exit(ex)
}

func Test_Chmod(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "chmod"
	test_fs.name = "test_a"
	test_fs.mode = os.FileMode(0755)
	test_fs.err = nil

	//err nil test
	err := test_client.Chmod(test_fs.name, test_fs.mode)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err != test_fs.err {
		t.Errorf("Chomd err expect:%s, get:%s", test_fs.err, err)
	}

	//err msg test
	test_fs.err = errors.New("test text1")

	err = test_client.Chmod(test_fs.name, test_fs.mode)

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Chomd err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_Chtimes(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "chtimes"
	test_fs.name = "test_b"
	test_fs.atime = time.Unix(1429348849, 0)
	test_fs.mtime = time.Unix(1429348849, 0)
	test_fs.err = errors.New("test text2")

	err := test_client.Chtimes(test_fs.name, test_fs.atime, test_fs.mtime)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Chtimes err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_Mkdir(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "mkdir"
	test_fs.name = "test_c"
	test_fs.perm = 0755
	test_fs.err = errors.New("test text3")

	err := test_client.Mkdir(test_fs.name, test_fs.perm)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Mkdir err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_MkdirAll(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "mkdirall"
	test_fs.name = "test_d"
	test_fs.perm = 0755
	test_fs.err = errors.New("test text4")

	err := test_client.MkdirAll(test_fs.name, test_fs.perm)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("MkdirAll err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_Remove(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "remove"
	test_fs.name = "test_e"
	test_fs.err = errors.New("test text5")

	err := test_client.Remove(test_fs.name)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Remove err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_RemoveAll(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "removeall"
	test_fs.name = "test_f"
	test_fs.err = errors.New("test text6")

	err := test_client.RemoveAll(test_fs.name)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("RemoveAll err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_Rename(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "rename"
	test_fs.oldpath = "test_a"
	test_fs.newpath = "test_b"
	test_fs.err = errors.New("test text7")

	err := test_client.Rename(test_fs.oldpath, test_fs.newpath)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Rename err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_Truncate(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "truncate"
	test_fs.name = "test_g"
	test_fs.size = 1235
	test_fs.err = errors.New("test text8")

	err := test_client.Truncate(test_fs.name, test_fs.size)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Truncate err expect:%s, get:%s", test_fs.err, err)
	}
}

func Test_Stat(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "stat"
	test_fs.name = "test_h"

	test_fs.fi = new(FileInfo)
	test_fs.fi.name = "test_a"
	test_fs.fi.size = 12345
	test_fs.fi.mode = 0766
	test_fs.fi.modtime = time.Unix(1429348849, 0)

	test_fs.err = errors.New("test text9")

	fi, err := test_client.Stat(test_fs.name)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Stat err expect:%s, get:%s", test_fs.err, err)
	}

	if fi.Name() != test_fs.fi.Name() {
		t.Errorf("Stat fi.name expect:%s, get:%s", test_fs.fi.Name(), fi.Name())
	}

	if fi.Size() != test_fs.fi.Size() {
		t.Errorf("Stat fi.size expect:%d, get:%d", test_fs.fi.Size(), fi.Size())
	}

	if fi.Mode() != test_fs.fi.Mode() {
		t.Errorf("Stat fi.mode expect:%o, get:%o", test_fs.fi.Mode(), fi.Mode())
	}

	if !fi.ModTime().Equal(test_fs.fi.ModTime()) {
		t.Errorf("Stat mtime expect:%s give:%s",
			test_fs.fi.ModTime().Format(time.RFC3339Nano),
			fi.ModTime().Format(time.RFC3339Nano))
	}
}

func Test_Lstat(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "lstat"
	test_fs.name = "test_i"

	test_fs.fi = new(FileInfo)
	test_fs.fi.name = "test_j"
	test_fs.fi.size = 12345
	test_fs.fi.mode = 0766
	test_fs.fi.modtime = time.Unix(1429348849, 0)

	test_fs.err = errors.New("test text10")

	fi, err := test_client.Lstat(test_fs.name)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Lstat err expect:%s, get:%s", test_fs.err, err)
	}

	if fi.Name() != test_fs.fi.Name() {
		t.Errorf("Lstat fi.name expect:%s, get:%s", test_fs.fi.Name(), fi.Name())
	}

	if fi.Size() != test_fs.fi.Size() {
		t.Errorf("Lstat fi.size expect:%d, get:%d", test_fs.fi.Size(), fi.Size())
	}

	if fi.Mode() != test_fs.fi.Mode() {
		t.Errorf("Lstat fi.mode expect:%o, get:%o", test_fs.fi.Mode(), fi.Mode())
	}

	if !fi.ModTime().Equal(test_fs.fi.ModTime()) {
		t.Errorf("Lstat fi.modtime expect:%s give:%s",
			test_fs.fi.ModTime().Format(time.RFC3339Nano),
			fi.ModTime().Format(time.RFC3339Nano))
	}
}

func Test_Create(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "create"
	test_fs.name = "test_k"
	test_fs.err = errors.New("test text11")
	test_fs.file = nil

	fd, err := test_client.Create(test_fs.name)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Create err expect:%s, get:%s", test_fs.err, err)
	}

	if fd != nil {
		t.Errorf("Create fd expect:nil, get:%v", fd)
	}
}

func Test_Open(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "open"
	test_fs.name = "test_l"
	test_fs.err = errors.New("test text12")
	test_fs.file = nil

	fd, err := test_client.Open(test_fs.name)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err.Error() != test_fs.err.Error() {
		t.Errorf("Open err expect:%s, get:%s", test_fs.err, err)
	}

	if fd != nil {
		t.Errorf("Open fd expect:nil, get:%v", fd)
	}
}

func Test_OpenFile(t *testing.T) {
	test_fs.t = t
	test_fs.expect_call = "openfile"
	test_fs.name = "test_m"
	test_fs.flag = 1
	test_fs.perm = 0711
	test_fs.err = nil

	test_file = new(testFile)
	test_fs.file = test_file

	fd, err := test_client.OpenFile(test_fs.name, test_fs.flag, test_fs.perm)

	if test_fs.call != test_fs.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_fs.expect_call, test_fs.call)
	}

	if err != nil {
		t.Errorf("Open err expect:nil, get:%v", err)
	}

	if fd == nil {
		t.Error("Open fd expect:file, get:nil")
		t.FailNow()
	}

	test_client_file = fd
}

func Test_f_Chmod(t *testing.T) {
	test_file.t = t
	test_file.mode = 0755
	test_file.expect_call = "chmod"
	test_file.err = errors.New("test text12")

	err := test_client_file.Chmod(test_file.mode)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err.Error() != test_file.err.Error() {
		t.Errorf("Open err expect:%s, get:%s", test_file.err, err)
	}
}

func Test_f_Read(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "read"
	test_file.b = []byte("test")

	b := make([]byte, 5)

	n, err := test_client_file.Read(b)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err != nil {
		t.Errorf("Read expect:nil, get:%v", err)
	}

	if string(b[:n]) != string(test_file.b) {
		t.Errorf("Read expect:%s, get:%s", string(test_file.b), string(b[:n]))
	}
}

func Test_f_ReadAt(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "readat"
	test_file.b = []byte("test")
	test_file.off = 1

	b := make([]byte, 5)

	n, err := test_client_file.ReadAt(b, test_file.off)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err != io.EOF {
		t.Errorf("Read expect:io.EOF, get:%v", err)
	}

	if string(b[:n]) != string(test_file.b[test_file.off:]) {
		t.Errorf("Read expect:%s, get:%s", string(test_file.b[test_file.off:]), string(b[:n]))
	}
}

func Test_f_Readdirnames(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "readdirnames"
	test_file.dirs = []string{"test"}
	test_file.n = 1
	test_file.err = nil

	dirs, err := test_client_file.Readdirnames(test_file.n)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err != nil {
		t.Errorf("Readdirnames expect:nil, get:%v", err)
	}

	if len(test_file.dirs) != len(dirs) || dirs[0] != test_file.dirs[0] {
		t.Errorf("Readdirnames expect:%v, get:%v", test_file.dirs, dirs)
	}
}

func Test_f_Seek(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "seek"
	test_file.offset = 12
	test_file.whence = 13
	test_file.ret = 2
	test_file.err = nil

	ret, err := test_client_file.Seek(test_file.offset, test_file.whence)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err != nil {
		t.Errorf("Seek expect:nil, get:%v", err)
	}

	if ret != test_file.ret {
		t.Errorf("Seek expect:%d, get:%d", test_file.ret, ret)
	}
}

func Test_f_Stat(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "stat"
	test_file.err = errors.New("test text12")

	_fi := new(FileInfo)
	_fi.name = "test_a"
	_fi.size = 12345
	_fi.mode = 0766
	_fi.modtime = time.Unix(1429348849, 0)
	test_file.fi = _fi

	fi, err := test_client_file.Stat()

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err.Error() != test_file.err.Error() {
		t.Errorf("Stat expect:%s, get:%s", test_file.err, err)
	}

	if fi.Name() != test_file.fi.Name() {
		t.Errorf("Stat fi.name expect:%s, get:%s", test_file.fi.Name(), fi.Name())
	}

	if fi.Size() != test_file.fi.Size() {
		t.Errorf("Stat fi.size expect:%d, get:%d", test_file.fi.Size(), fi.Size())
	}

	if fi.Mode() != test_file.fi.Mode() {
		t.Errorf("Stat fi.mode expect:%o, get:%o", test_file.fi.Mode(), fi.Mode())
	}

	if !fi.ModTime().Equal(test_file.fi.ModTime()) {
		t.Errorf("Stat mtime expect:%s give:%s",
			test_file.fi.ModTime().Format(time.RFC3339Nano),
			fi.ModTime().Format(time.RFC3339Nano))
	}
}

func Test_f_sync(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "sync"
	test_file.err = errors.New("test text14")

	err := test_client_file.Sync()

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err.Error() != test_file.err.Error() {
		t.Errorf("Stat expect:%s, get:%s", test_file.err, err)
	}
}

func Test_f_Truncate(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "truncate"
	test_file.size = 123
	test_file.err = errors.New("test text15")

	err := test_client_file.Truncate(test_file.size)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err.Error() != test_file.err.Error() {
		t.Errorf("Stat expect:%s, get:%s", test_file.err, err)
	}
}

func Test_f_Write(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "write"
	test_file.b = []byte("1230")
	test_file.n = 3
	test_file.err = errors.New("test text15")

	n, err := test_client_file.Write(test_file.b)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err.Error() != test_file.err.Error() {
		t.Errorf("Stat expect:%s, get:%s", test_file.err, err)
	}

	if n != test_file.n {
		t.Errorf("Write expect:%n, get:%n", test_file.n, n)
	}
}

func Test_f_WriteAt(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "writeat"
	test_file.b = []byte("1230")
	test_file.off = 1
	test_file.n = 3
	test_file.err = errors.New("test text15")

	n, err := test_client_file.WriteAt(test_file.b, test_file.off)

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err.Error() != test_file.err.Error() {
		t.Errorf("WriteAt expect:%s, get:%s", test_file.err, err)
	}

	if n != test_file.n {
		t.Errorf("WriteAt expect:%n, get:%n", test_file.n, n)
	}
}

func Test_f_Close(t *testing.T) {
	test_file.t = t
	test_file.expect_call = "close"
	test_file.err = errors.New("test text12")

	err := test_client_file.Close()

	if test_file.call != test_file.expect_call {
		t.Errorf("Expect Call :%s, get:%s", test_file.expect_call, test_file.call)
	}

	if err.Error() != test_file.err.Error() {
		t.Errorf("Open err expect:%s, get:%s", test_file.err, err)
	}
}

func Test_Close(t *testing.T) {
	err := test_client.Close()

	if err != nil {
		t.Errorf("Close expect:nil, get:%v", err)
	}
}
