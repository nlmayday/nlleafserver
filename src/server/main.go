package main

import (
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"

	nlleaf "github.com/nlmayday/nlleaf"
	lconf "github.com/nlmayday/nlleaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	nlleaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
