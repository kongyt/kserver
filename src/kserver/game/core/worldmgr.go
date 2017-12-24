package core

import (
	"sync"
	"github.com/kongyt/leaf/gate"
	"errors"
	"kserver/msg"
	"github.com/golang/protobuf/proto"
)

type WorldMgr struct{
	PlayerNumGen	int32
	Players 		map[int32]*Player
	AoiObj1			*AOIMgr
	sync.RWMutex
}

var WorldMgrObj *WorldMgr

func init(){
	WorldMgrObj = &WorldMgr{
		PlayerNumGen:	0,
		Players: 		make(map[int32]*Player),
	}
}

func (this *WorldMgr)AddPlayer(agent gate.Agent)(*Player, error){
	this.Lock()
	this.PlayerNumGen += 1
	p := NewPlayer(agent, this.PlayerNumGen)
	this.Players[p.Pid] = p
	this.Unlock()

	res := &msg.S2C_Enter_World_Res{
		Result: proto.Bool(true),
		Pid: proto.Int32(p.Pid),
	}
	p.Agent.WriteMsg(res)

	this.AoiObj1.Add2AOI(p)

	p.SyncSurrounding()

	return p, nil

	return p, nil
}

func (this *WorldMgr)RemovePlayer(pid int32){
	this.Lock()
	defer this.Unlock()

	this.AoiObj1.LeaveAOI(this.Players[pid])

	delete(this.Players, pid)
}

func (this *WorldMgr)MovePlayer(p *Player){
	playerMoveNoti := &msg.S2C_Player_Move_And_Action_Noti{
		Pid:  proto.Int32(p.Pid),
		Pos: &msg.Position{
			X:	proto.Float32(p.X),
			Y:  proto.Float32(p.Y),
			Z:  proto.Float32(p.Z),
			V:  proto.Float32(p.V),
		},
		Action: proto.Int32(0),
	}

	pids, err := this.AoiObj1.GetSurroundingPids(p)
	if err == nil{
		for _, pid := range  pids{
			player, err1 := this.GetPlayer(pid)
			if err1 == nil{
				player.Agent.WriteMsg(playerMoveNoti)
			}
		}
	}

}

func (this *WorldMgr)GetPlayer(pid int32)(*Player, error){
	this.RLock()
	defer this.RUnlock()
	p, ok := this.Players[pid]
	if ok{
		return p, nil
	}else{
		return nil, errors.New("no play in the world")
	}
}

func (this *WorldMgr) BroadcastToAll(msg interface{}){
	this.RLock()
	defer this.RUnlock()
	for _, p := range this.Players{
		p.Agent.WriteMsg(msg)
	}
}

