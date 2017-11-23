package base

import (
	"github.com/kongyt/leaf/module"
	"kserver/conf"
	"github.com/kongyt/leaf/chanrpc"
)

func NewSkeleton()*module.Skeleton{
	skeleton := &module.Skeleton{
		GoLen:				conf.GoLen,
		TimerDispatcherLen:	conf.TimerDispatcherLen,
		AsynCallLen:		conf.AsynCallLen,
		ChanRPCServer:		chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}