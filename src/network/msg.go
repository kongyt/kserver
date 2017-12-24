package network

import (
	"github.com/kongyt/leaf/network"
	"github.com/golang/protobuf/proto"
	"sync"
	"errors"
	"fmt"
	"github.com/kongyt/leaf/log"
)

type MsgHandler func(conn network.Conn, message proto.Message)

type MsgHandlerInfo struct{
	msgId 	int32
	message proto.Message
	handler MsgHandler
}

type IMsgProcesser interface{
	Register(msgId int32, message proto.Message, handler MsgHandler) error
	Handle(msgId int32, data []byte)
}

type MsgProcesser struct{
	sync.RWMutex
	handlers	map[int32] *MsgHandlerInfo
}

func (this *MsgProcesser)Register(msgId int32, message proto.Message, handler MsgHandler) error{
	this.Lock()
	defer this.Unlock()

	if _, ok := this.handlers[msgId]; ok == true{
		return errors.New(fmt.Sprintf("msgId(%d) already registered", msgId))
	}

	info := &MsgHandlerInfo{
		msgId: 	 msgId,
		message: message,
		handler: handler,
	}
	this.handlers[msgId] = info

	return nil
}

func (this *MsgProcesser)Handle(msgId int32, data []byte){
	this.Lock()
	defer this.Unlock()
	info, ok := this.handlers[msgId]
	if ok != true{
		log.Debug("msgId(%d) not register", msgId)
	}else{
		proto.Unmarshal(data, info.message)
		info.handler(nil, info.message)
	}


}

