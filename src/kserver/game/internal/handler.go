package internal

import (
	"kserver/msg"
	"reflect"
	"github.com/kongyt/leaf/gate"
	"kserver/game/core"
	"github.com/golang/protobuf/proto"
)

func handleMsg(m interface{}, h interface{}){
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init(){
	handleMsg(&msg.C2S_Enter_World_Req{}, onEnterWorldReq)
	handleMsg(&msg.C2S_Sync_Player_Req{}, onSyncPlayerReq)
	handleMsg(&msg.C2S_Move_And_Action_Req{}, onMoveAndActionReq)
	handleMsg(&msg.C2S_Talk_Req{}, onTalkReq)
}

func onEnterWorldReq(args[] interface{}){
	// 收到登陆消息
	m := args[0].(*msg.C2S_Enter_World_Req)
	a := args[1].(gate.Agent)

	_ = m

	player, err := core.WorldMgrObj.AddPlayer(a)
	if err != nil{
		panic(err)
	}
	a.SetUserData(player.Pid)

}

func onSyncPlayerReq(args[] interface{}){
	m := args[0].(*msg.C2S_Sync_Player_Req)
	a := args[1].(gate.Agent)

	player, err := core.WorldMgrObj.GetPlayer(m.GetPid())
	if err != nil {
		panic(err)
	}

	res := &msg.S2C_Sync_Player_Res{
		Player: &msg.Player{
			Pid: proto.Int32(player.Pid),
			Pos: &msg.Position{
				X:	proto.Float32(player.X),
				Y:  proto.Float32(player.Y),
				Z:  proto.Float32(player.Z),
				V:  proto.Float32(player.V),
			},
		},
	}
	a.WriteMsg(res)


}

func onMoveAndActionReq(args[] interface{}){
	m := args[0].(*msg.C2S_Move_And_Action_Req)
	a := args[1].(gate.Agent)

	player, err := core.WorldMgrObj.GetPlayer(a.UserData().(int32))
	if err != nil {
		panic(err)
	}

	player.UpdatePos(m.Pos.GetX(), m.Pos.GetY(), m.Pos.GetZ(), m.Pos.GetV())
}

func onTalkReq(args[] interface{}){
	m := args[0].(*msg.C2S_Talk_Req)
	a := args[1].(gate.Agent)

	player, err := core.WorldMgrObj.GetPlayer(a.UserData().(int32))
	if err != nil {
		panic(err)
	}

	player.Talk(m.GetContext())
}
