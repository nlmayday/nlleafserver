package game

import (
	"server/game/internal"
	"server/global"
	"server/model"
)

var (
	Module    = new(internal.Module)
	ChanRPC   = internal.ChanRPC
	PlayerMgr = internal.NewPlayerMgr()
	Game      = internal.NewGame()
)

func InitGame() {
	var rows []*model.GameServers
	global.G_DB.Where("status = ?", 1).Find(&rows)
	for _, row := range rows {
		gs := Game.LoadGameServers(row)
		go gs.Run()
	}
}
