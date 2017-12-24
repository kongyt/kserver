package internal

import (
	"reflect"
	"kserver/msg"
	"github.com/golang/protobuf/proto"
	"github.com/kongyt/leaf/gate"
	"github.com/kongyt/leaf/log"
)

func handleMsg(m interface{}, h interface{}){
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init(){
	handleMsg(&msg.C2S_Register_Req{}, onRegisterReq)
	handleMsg(&msg.C2S_Login_Req{}, onLoginReq)
}

func onRegisterReq(args[] interface{}){
	// 收到注册消息
	m := args[0].(*msg.C2S_Register_Req)
	a := args[1].(gate.Agent)

	log.Debug("onRegisterReq(%v, %v)", m.GetUserName(), m.GetPassword())

	a.WriteMsg(&msg.S2C_Register_Res{
		Result: proto.Bool(true),
	})
}

func onLoginReq(args[] interface{}){
	// 收到登陆消息
	m := args[0].(*msg.C2S_Login_Req)
	a := args[1].(gate.Agent)

	log.Debug("onLoginReq(%v, %v)",m.GetUserName(), m.GetPassword())

	a.WriteMsg(&msg.S2C_Login_Res{
		Result: proto.Bool(true),
	})


}
