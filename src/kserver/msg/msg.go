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
}
