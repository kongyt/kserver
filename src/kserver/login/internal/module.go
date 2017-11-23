package internal

import (
	"kserver/base"
	"github.com/kongyt/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC = skeleton.ChanRPCServer
)

type Module struct{
	*module.Skeleton
}

func (m *Module) OnInit(){
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy(){

}
