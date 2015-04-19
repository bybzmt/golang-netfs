package netfs

import (
	"os"
	"io"
	"net"
	"errors"
	"bufio"
	"fmt"
	"time"
	"encoding/binary"
)

type conn struct{
	conn net.Conn
	buf *bufio.ReadWriter
}

func (c *conn) init(conn net.Conn, rBuf, wBuf int) {
	c.conn = conn
	r := bufio.NewReaderSize(conn, rBuf)
	w := bufio.NewWriterSize(conn, wBuf)
	c.buf = bufio.NewReadWriter(r, w)
}

func (c *conn) LinkInit() (err error) {
	defer onPanic(&err)

	c.writeUint8(TYTE_INIT)
	c.writeUint32(VERSION)
	c.flush()

	t := c.readUint8()
	if t != TYTE_INIT {
		return errors.New("Protocol Unexpect.")
	}

	ver := c.readUint32()
	if ver != VERSION {
		return errors.New("Protocol Version Unexpect.")
	}

	return nil
}

func (c *conn) flush() {
	err := c.buf.Flush()
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) setDeadline(t time.Duration) {
	err := c.conn.SetDeadline(time.Now().Add(t))
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) Close() {
	err := c.conn.Close()
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

// ----- 请求/响应 -----

func (c *conn) doRequest(code uint8) {
	c.writeUint8(TYPE_REQUEST)
	c.writeUint8(code)
}

func (c *conn) waitResponse(code uint8) {
	c.flush()

	_type := c.readUint8()
	if _type != TYPE_RESPONSE {
		panic(Data_Error(fmt.Sprintf("Expect Response(%d) Get:%d", TYPE_RESPONSE, _type)))
	}

	_code := c.readUint8()
	if _code != code {
		panic(Data_Error(fmt.Sprintf("Expect Target:%d, Get:%d", code, _code)))
	}
}

//------------

func (c *conn) doResponse(code uint8) {
	c.writeUint8(TYPE_RESPONSE)
	c.writeUint8(code)
}

func (c *conn) waitRequest() uint8 {
	c.flush()

	_type := c.readUint8()
	if _type != TYPE_REQUEST {
		panic(Data_Error(fmt.Sprintf("Expect Request(%d) Get:%d", TYPE_REQUEST, _type)))
	}

	return c.readUint8()
}


// ----- uint 读取 -----

func (c *conn) readUint8() (number uint8) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

func (c *conn) readUint16() (number uint16) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

func (c *conn) readUint32() (number uint32) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

func (c *conn) readUint64() (number uint64) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

// ----- int 读取 -----

func (c *conn) readInt8() (number int8) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

func (c *conn) readInt16() (number int16) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

func (c *conn) readInt32() (number int32) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

func (c *conn) readInt64() (number int64) {
	err := binary.Read(c.buf, binary.BigEndian, &number)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
	return
}

// ------ uint 写入 -----------

func (c *conn) writeUint8(data uint8) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) writeUint16(data uint16) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) writeUint32(data uint32) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) writeUint64(data uint64) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

// ------ uint 写入 -----------

func (c *conn) writeInt8(data int8) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) writeInt16(data int16) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) writeInt32(data int32) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) writeInt64(data int64) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

// ----- 字符串 读取 -----

func (c *conn) readFull(b []byte) {
	_, err := io.ReadFull(c.buf, b)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}

func (c *conn) readString() string {
	_len32 := c.readUint32()

	b := make([]byte, int(_len32))
	c.readFull(b)

	return string(b)
}

func (c *conn) readByte() []byte {
	_len32 := c.readUint32()

	b := make([]byte, int(_len32))
	c.readFull(b)

	return b
}

func (c *conn) readByteTo(b []byte) int {
	_len32 := c.readUint32()
	_len := int(_len32)

	if _len > len(b) {
		panic(IO_Error("Read c.buf too Small"))
	}

	c.readFull(b[:_len])
	return _len
}

func (c *conn) readError() error {
	_len16 := c.readUint16()

	if _len16 > ERROR_MAX {
		switch _len16 {
		case ERROR_NIL : return nil
		case ERROR_EOF : return io.EOF
		default:
			panic(IO_Error("ReadError Len Not Defined"))
		}
	}

	_len := int(_len16)
	b := make([]byte, _len)

	c.readFull(b)

	return errors.New(string(b))
}


// ------ 字符串 写入 ------------

func (c *conn) writeString(s string) {
	c.writeUint32(uint32(len(s)))
	c.writeData([]byte(s))
}

func (c *conn) writeByte(s []byte) {
	c.writeUint32(uint32(len(s)))
	c.writeData(s)
}

func (c *conn) writeError(err error) {
	switch err {
	case nil :
		c.writeUint16(ERROR_NIL)
		return
	case io.EOF :
		c.writeUint16(ERROR_EOF)
		return
	}

	errs := err.Error()

	if len(errs) > int(ERROR_MAX) {
		errs = errs[:int(ERROR_MAX)]
	}

	c.writeUint16(uint16(len(errs)))
	c.writeData([]byte(errs))
}

func (c *conn) writeData(data interface{}) {
	err := binary.Write(c.buf, binary.BigEndian, data)
	if err != nil {
		panic(IO_Error(err.Error()))
	}
}


//----------------

func (c *conn) writeFileInfo(fi os.FileInfo) {
	c.writeString(fi.Name())
	c.writeInt64(fi.Size())
	c.writeUint32(uint32(fi.Mode()))
	c.writeInt64(fi.ModTime().Unix())
}

func (c *conn) readFileInfo() os.FileInfo {
	fi := new(FileInfo)
	fi.name = c.readString()
	fi.size = c.readInt64()
	fi.mode = os.FileMode(c.readUint32())
	fi.modtime = time.Unix(c.readInt64(), 0)
	return fi
}
