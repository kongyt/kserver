package network

import (
	"net"
	"encoding/binary"
	"sync"
)


type Conn interface {
	ReadMsg() ([]byte, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}

type TcpConn struct{
	Conn 		net.Conn
	UserData 	interface{}
	closeFlag	bool
}

func (this *TcpConn)ReadMsg()([]byte, error){
	return nil, nil
}

func (this *TcpConn)WriteRawMsg(data []byte) error{
	_, err := this.Conn.Write(data)
	return err
}

func (this *TcpConn)WriteMsg(msgId int32, data []byte)error{
	dataLen := 6 + len(data)
	buf := make([]byte, dataLen)
	binary.BigEndian.PutUint16(buf, uint16(dataLen))
	binary.BigEndian.PutUint32(buf[2:], uint32(msgId))
	copy(buf[6:], data)
	_, err := this.Conn.Write(buf)

	return err
}

func (this *TcpConn)LocalAddr() net.Addr{
	return this.Conn.LocalAddr()
}

func (this *TcpConn)RemoteAddr() net.Addr{
	return this.Conn.RemoteAddr()
}

func (this *TcpConn)Close(){
	this.Conn.Close()
}

func (this *TcpConn)Destroy(){
	this.Conn.
}