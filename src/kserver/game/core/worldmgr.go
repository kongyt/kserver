package core

import (
	"sync"
	"github.com/kongyt/leaf/gate"
	"errors"
)

type WorldMgr struct{
	PlayerNumGen	int32
	Players 		map[int32]*Player
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


	//TODU

	return p, nil
}

func (this *WorldMgr)RemovePlayer(pid int32){
	this.Lock()
	defer this.Unlock()
	delete(this.Players, pid)
}

func (this *WorldMgr)MovePlayer(p *Player){

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

