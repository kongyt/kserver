package network

import "net"

type INetFilter interface{
	Filter(ip net.IP) bool		// 返回false则表示过滤掉，直接断开连接
}

type DefaultNetFilter struct{

}

func (this *DefaultNetFilter)Filter(ip net.IP) bool{
	return true
}

