package core

import (
	"math/rand"
	"github.com/kongyt/leaf/gate"
	"kserver/msg"
	"github.com/golang/protobuf/proto"
	"errors"
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
	pids, err := WorldMgrObj.AoiObj1.GetSurroundingPids(this)

	if err == nil {
		// 同步周围的人
		syncPlayersNoti := &msg.S2C_Sync_Players_Noti{}

		// 同步自己到周围的人
		syncPlayerNoti := &msg.S2C_Sync_Player_Noti{
			Player: &msg.Player{
				Pid: proto.Int32(this.Pid),
				Pos: &msg.Position{
					X:	proto.Float32(this.X),
					Y:  proto.Float32(this.Y),
					Z:  proto.Float32(this.Z),
					V:  proto.Float32(this.V),
				},
			},
		}
		for _, pid := range pids{
			player, err1 := WorldMgrObj.GetPlayer(pid)
			if err1 == nil{
				p := &msg.Player{
					Pid:	proto.Int32(player.Pid),
					Pos:	&msg.Position{
						X:	proto.Float32(player.X),
						Y:	proto.Float32(player.Y),
						Z:	proto.Float32(player.Z),
						V:  proto.Float32(player.V),
					},
				}
				syncPlayersNoti.Players = append(syncPlayersNoti.Players, p)
				player.Agent.WriteMsg(syncPlayerNoti)
			}
		}

		// 分包发送
		per := 20
		ps := syncPlayersNoti.Players
		for i:=0; ;i++{
			if i*per > len(ps) - 1{
				break;
			}
			if i *per + per > len(ps) - 1{
				syncPlayersNoti.Players = ps[i*per:]
			}else{
				syncPlayersNoti.Players = ps[i*per: i*per+per]
			}
			this.Agent.WriteMsg(syncPlayersNoti)
		}
	}else{
		panic(err)
	}
}

// 更新位置
func (this *Player) UpdatePos(x float32, y float32, z float32, v float32){
	oldGridId := WorldMgrObj.AoiObj1.GetGridIDByPos(this.X, this.Z)

	// 更新位置的时候判断是否需要更新GridID
	newGridId := WorldMgrObj.AoiObj1.GetGridIDByPos(x, z)

	if newGridId < 0 || newGridId >= WorldMgrObj.AoiObj1.lenX * WorldMgrObj.AoiObj1.lenY{
		return
	}

	// 更新
	this.X = x
	this.Y = y
	this.Z = z
	this.V = v

	if oldGridId != newGridId{
		WorldMgrObj.AoiObj1.LeaveAOIFromGrid(this, oldGridId)
		WorldMgrObj.AoiObj1.Add2AOI(this)

		// 需要处理老的AOI消失和新的AOI出生
		this.OnExchangeAoiGrid(oldGridId, newGridId)
	}

	WorldMgrObj.MovePlayer(this)

}

func (this *Player)OnExchangeAoiGrid(oldGridId int32, newGridId int32) error{
	oldAoiGrids, err1 := WorldMgrObj.AoiObj1.GetSurroundingByGridId(oldGridId)
	newAoiGrids, err2 := WorldMgrObj.AoiObj1.GetSurroundingByGridId(newGridId)
	if err1 != nil || err2 != nil{
		return errors.New("OnExchangeAoiGrid")
	}
	alls := make([]*Grid, 0)
	alls = append(alls, oldAoiGrids...)
	alls = append(alls, newAoiGrids...)
	//并集
	union := make(map[int32]*Grid, 0)
	for _, v:= range alls{
		if _, ok := union[v.ID]; ok != true{
			union[v.ID] = v
		}
	}
	oldAoiGridsMap := make(map[int32]bool, 0)
	for _, oldGrid := range oldAoiGrids{
		if _, ok := oldAoiGridsMap[oldGrid.ID]; ok != true{
			oldAoiGridsMap[oldGrid.ID] = true
		}
	}

	newAoiGridsMap := make(map[int32]bool, 0)
	for _, newGrid := range newAoiGrids{
		if _, ok := newAoiGridsMap[newGrid.ID]; ok != true{
			newAoiGridsMap[newGrid.ID] = true
		}
	}

	for gid, grid := range union{
		//出生
		if _, ok := oldAoiGridsMap[gid]; ok != true{

			// 同步周围的人
			syncPlayersNoti := &msg.S2C_Sync_Players_Noti{}

			// 同步自己到周围的人
			syncPlayerNoti := &msg.S2C_Sync_Player_Noti{
				Player: &msg.Player{
					Pid: proto.Int32(this.Pid),
					Pos: &msg.Position{
						X:	proto.Float32(this.X),
						Y:  proto.Float32(this.Y),
						Z:  proto.Float32(this.Z),
						V:  proto.Float32(this.V),
					},
				},
			}

			for _, pid := range grid.GetPids(){
				if pid != this.Pid{
					player, err := WorldMgrObj.GetPlayer(pid)
					if err == nil{
						p := &msg.Player{
							Pid:	proto.Int32(player.Pid),
							Pos:	&msg.Position{
								X:	proto.Float32(player.X),
								Y:	proto.Float32(player.Y),
								Z:	proto.Float32(player.Z),
								V:  proto.Float32(player.V),
							},
						}
						syncPlayersNoti.Players = append(syncPlayersNoti.Players, p)
						player.Agent.WriteMsg(syncPlayerNoti)
					}
				}
			}
			// 分包发送
			per := 20
			ps := syncPlayersNoti.Players
			for i:=0; ;i++{
				if i*per > len(ps) - 1{
					break;
				}
				if i *per + per > len(ps) - 1{
					syncPlayersNoti.Players = ps[i*per:]
				}else{
					syncPlayersNoti.Players = ps[i*per: i*per+per]
				}
				this.Agent.WriteMsg(syncPlayersNoti)
			}
		}

		if _, ok := newAoiGridsMap[gid]; ok != true{
			//消失
			playersLeaveNoti := &msg.S2C_Players_Leave_View_Noti{}
			playerLeaveNoti := &msg.S2C_Player_Leave_View_Noti{
				Pid: proto.Int32(this.Pid),
			}
			for _, pid := range grid.GetPids(){
				if pid != this.Pid{
					player, err := WorldMgrObj.GetPlayer(pid)
					if err == nil{
						playersLeaveNoti.Pids = append(playersLeaveNoti.Pids, player.Pid)
						player.Agent.WriteMsg(playerLeaveNoti)
					}
				}
			}
			this.Agent.WriteMsg(playersLeaveNoti)
		}
	}
	return nil
}


func (this *Player)Talk(context string){
	talkNoti := &msg.S2C_Player_Talk_Noti{
		Context: proto.String(context),
	}
	WorldMgrObj.BroadcastToAll(talkNoti)
}

func (this *Player)LostConnection()  {
	playerLeaveNoti := msg.S2C_Player_Leave_View_Noti{
		Pid: proto.Int32(this.Pid),
	}
	WorldMgrObj.BroadcastToAll(playerLeaveNoti)
}


