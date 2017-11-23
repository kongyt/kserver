package main

import (
	"github.com/kongyt/leaf"
	lconf "github.com/kongyt/leaf/conf"
	"kserver/conf"
	"kserver/game"
	"kserver/gate"
	"kserver/login"
)


func main(){
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}

