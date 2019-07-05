package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:5], pkgLen)
	n, err := this.Conn.Write(this.Buf[0:5])
	if n != 5 || err != nil {
		fmt.Println("write err=", err)
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("write err=", err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	n, err := this.Conn.Read(this.Buf[:5])
	if n != 5 || err != nil {
		fmt.Println("Read err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:5])
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}
