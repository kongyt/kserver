package msg

import (
	"github.com/kongyt/leaf/network/protobuf"
	"github.com/golang/protobuf/proto"
)

var Processor = protobuf.NewProcessor()

func register(msgId EMsg, msg proto.Message){
	Processor.Register(uint32(msgId), msg)
}

func init(){
	register(EMsg_C2S_Register_Req_ID, &C2S_Register_Req{})
	register(EMsg_S2C_Register_Res_ID, &S2C_Register_Res{})

	register(EMsg_C2S_Login_Req_ID, &C2S_Login_Req{})
	register(EMsg_S2C_Login_Res_ID, &S2C_Login_Res{})

	register(EMsg_C2S_Enter_World_Req_ID, &C2S_Enter_World_Req{})
	register(EMsg_S2C_Enter_World_Res_ID, &S2C_Enter_World_Res{})

	register(EMsg_C2S_Move_And_Action_Req_ID, &C2S_Move_And_Action_Req{})
	register(EMsg_S2C_Move_And_Action_Res_ID, &S2C_Move_And_Action_Res{})

	register(EMsg_C2S_Sync_Player_Req_ID, &C2S_Sync_Player_Req{})
	register(EMsg_S2C_Sync_Player_Res_ID, &S2C_Sync_Player_Res{})

	register(EMsg_C2S_Talk_Req_ID, &C2S_Talk_Req{})
	register(EMsg_S2C_Talk_Res_ID, &S2C_Talk_Res{})

	register(EMsg_S2C_Player_Leave_View_Noti_ID, &S2C_Player_Leave_View_Noti{})
	register(EMsg_S2C_Players_Leave_View_Noti_ID, &S2C_Players_Leave_View_Noti{})
	register(EMsg_S2C_Player_Talk_Noti_ID, &S2C_Player_Talk_Noti{})
	register(EMsg_S2C_Sync_Players_Noti_ID, &S2C_Sync_Players_Noti{})
	register(EMsg_S2C_Sync_Player_Noti_ID, &S2C_Sync_Player_Noti{})
	register(EMsg_S2C_Player_Move_And_Action_Noti_ID, &S2C_Player_Move_And_Action_Noti{})
}
