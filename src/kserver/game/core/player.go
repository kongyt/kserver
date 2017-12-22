package core

import (
	"math/rand"
	"github.com/golang/protobuf/proto"
	"github.com/kongyt/leaf/gate"
)

type Player struct {
	Agent	gate.Agent
	Pid		int32
	X 		float32 // 平面X
	Y       float32 // 高度
	Z 		float32 // 平面y
	V 		float32 // 旋转0-360度
}

func NewPlayer(agent gate.Agent, pid int32)*Player{
	p := &Player{
		Agent:	agent,
		Pid:	pid,
		X:		float32(rand.Intn(10)+160),
		Y:		0,
		Z:		float32(rand.Intn(17	)+134),
		V:		0,
	}
	return p
}

// 同步周围玩家
func (this *Player) SyncSurrounding(){

}

// 更新位置
func (this *Player) UpdatePos(x float32, y float32, z float32, v float32){

	// 更新
	this.X = x
	this.Y = y
	this.Z = z
	this.V = v
}

func (this *Player)Talk(context string){

}

func (this *Player)LostConnection()  {

}


