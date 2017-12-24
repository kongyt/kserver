package internal

import (
	"github.com/kongyt/leaf/gate"
	//"kserver/game/core"
)

func init(){
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	//player, err := core.WorldMgrObj.GetPlayer(a.UserData().(int32))
	//if err != nil {
	//	return
	//}

	//core.WorldMgrObj.RemovePlayer(player.Pid)
	//player.LostConnection()
}