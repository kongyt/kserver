package network

import (
	"net"
	"sync"
)

type IMsgParser interface{

}

type TcpServer struct{
	Addr            string
	MaxConnNum      int
	listener        net.Listener
	Conns 			map[net.Conn]struct{}
	mutexConns      sync.Mutex
	wgLn            sync.WaitGroup
	wgConns         sync.WaitGroup

	msgParser    	*IMsgParser

}
