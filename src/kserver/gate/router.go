package gate

import (
	"kserver/msg"
	"kserver/login"
)

func init(){
	msg.Processor.SetRouter(&msg.C2S_Register_Req{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Login_Req{}, login.ChanRPC)
}
