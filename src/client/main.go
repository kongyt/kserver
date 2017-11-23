package main

import (
	"encoding/binary"
	"net"
	"kserver/msg"
	"github.com/golang/protobuf/proto"
	"github.com/kongyt/leaf/log"
	"time"
)

type MsgProcesser struct{
	msgInfo      map[uint32] *MsgInfo
}

type MsgHandler func([]interface{})

type MsgInfo struct{
	msgHandler		MsgHandler
	msg				proto.Message
}

type Server struct{
	conn 		net.Conn
	runFlag		chan bool
	connected   bool
}

var Processer *MsgProcesser
var server *Server

func init(){
	Processer = new(MsgProcesser)
	Processer.msgInfo = make(map[uint32]*MsgInfo)

	register(msg.EMsg_S2C_Register_Res_ID, &msg.S2C_Register_Res{}, onRegisterRes)
	register(msg.EMsg_S2C_Login_Res_ID, &msg.S2C_Login_Res{}, onLoginRes)
}

func register(msgId msg.EMsg,  msg proto.Message, msgHandler MsgHandler){
	if _, ok := Processer.msgInfo[uint32(msgId)]; ok{
		log.Fatal("msg already registered")
	}

	msgInfo := new (MsgInfo)
	msgInfo.msg = msg
	msgInfo.msgHandler = msgHandler
	Processer.msgInfo[uint32(msgId)] = msgInfo
}



func run(conn net.Conn){
	recvBuf := make([]byte, 1024)
	bufDataLen := 0

	for {
		newDataLen, err := conn.Read(recvBuf[bufDataLen:])
		if err != nil{
			return
		}
		bufDataLen += newDataLen

		dataLen := binary.BigEndian.Uint16(recvBuf)
		if uint32(dataLen) >= uint32(bufDataLen) - 2{
			msgId := binary.BigEndian.Uint32(recvBuf[2:6])
			if msg, ok := Processer.msgInfo[msgId]; ok{
				proto.Unmarshal(recvBuf[6:], msg.msg)
				if msg.msgHandler != nil {
					msg.msgHandler([]interface{}{msg.msg, conn})
					copy(recvBuf[0:], recvBuf[2+dataLen:])
					bufDataLen -= (2+int(dataLen))
				}
			}

		}
	}
}

func sendMsg(conn net.Conn, msgId uint32, message proto.Message){

	data, err := proto.Marshal(message)

	if err != nil{
		return
	}

	m:= make([]byte, 2+4+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)+4))
	binary.BigEndian.PutUint32(m[2:], msgId)
	copy(m[6:], data)

	conn.Write(m)
}

func main(){
	server = new (Server)

	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil{
		panic(err)
	}

	server.conn = conn
	server.connected = true

	go run(conn)

	for {
		sendRegisterReq("kongyatong", "pppp")
		sendLoginReq("kong", "password")
		time.Sleep(time.Second)
	}

	time.Sleep(1000)

}

func sendRegisterReq(username string, password string){
	if server.connected == false{
		return
	}

	req := &msg.C2S_Register_Req{
		UserName: proto.String(username),
		Password: proto.String(password),
	}

	sendMsg(server.conn, uint32(msg.EMsg_C2S_Register_Req_ID), req)
}

func onRegisterRes(args[] interface{}){
	// 收到注册消息回复
	//m := args[0].(*msg.S2C_Register_Res)
	//a := args[1].(net.Conn)

	log.Debug("注册成功")
}


func sendLoginReq(username string, password string){
	if server.connected == false{
		return
	}

	req := &msg.C2S_Login_Req{
		UserName: proto.String(username),
		Password: proto.String(password),
	}

	sendMsg(server.conn, uint32(msg.EMsg_C2S_Login_Req_ID), req)
}

func onLoginRes(args[] interface{}){
	// 收到注册消息回复
	//m := args[0].(*msg.S2C_Register_Res)
	//a := args[1].(net.Conn)

	log.Debug("登陆成功")
}