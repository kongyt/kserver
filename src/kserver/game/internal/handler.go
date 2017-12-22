package internal

import (
	"kserver/msg"
	"reflect"
)

func handleMsg(m interface{}, h interface{}){
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init(){
	handleMsg(&msg.C2S_Enter_World_Req{}, onEnterWorldReq)
	handleMsg(&msg.C2S_Sync_Player_Req{}, onSyncPlayerReq)
	handleMsg(&msg.C2S_Move_Req{}, onMoveReq)
}

func onEnterWorldReq(args[] interface{}){

}

func onSyncPlayerReq(args[] interface{}){

}

func onMoveReq(args[] interface{}){

}
