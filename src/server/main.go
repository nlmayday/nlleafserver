package main

import (
	"log"
	"server/conf"
	"server/game"
	"server/gate"
	"server/global"
	"server/login"

	nlleaf "github.com/nlmayday/nlleaf"
	lconf "github.com/nlmayday/nlleaf/conf"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	global.InitGlobal()
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	// 加载游戏配置
	game.InitGame()
	nlleaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
