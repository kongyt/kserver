package gate

import (
	"kserver/msg"
	"kserver/login"
	"kserver/game"
)

func init(){
	msg.Processor.SetRouter(&msg.C2S_Register_Req{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Login_Req{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.C2S_Enter_World_Req{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Move_And_Action_Req{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Sync_Player_Req{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Talk_Req{}, game.ChanRPC)
}
