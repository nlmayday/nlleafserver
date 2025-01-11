package internal

import (
	"server/base"
	"server/global"
	"server/model"
	protojson "server/protocol/json"

	"go.uber.org/zap"
)

var _game *Game

type Game struct {
	GameServeMap map[uint64]*GameServe
}

func NewGame() *Game {
	_game = &Game{
		GameServeMap: make(map[uint64]*GameServe),
	}

	return _game
}

func (g *Game) GetGameServe(gameId uint64) *GameServe {
	return g.GameServeMap[gameId]
}
func (g *Game) LoadGameServers(game *model.GameServers) *GameServe {
	gs := g.GetGameServe(game.ID)
	if gs == nil {
		gs = new(GameServe)
		gs.LoadGame(game)
		g.GameServeMap[game.ID] = gs
		return gs
	} else {
		gs.Destroy()
		gs.LoadGame(game)
	}
	return gs
}

func (g *Game) JoinGameReq(userInfo *base.UserInfo, m *protojson.JoinGameReq) {
	gameServer := g.GetGameServe(m.GameId)
	if gameServer == nil {
		global.G_LOG.Error("游戏不存在", zap.Uint64("gameId", m.GameId))
		return
	}
	gameServer.JoinGameReq(userInfo, m)
}

func (g *Game) LeaveGameReq(userInfo *base.UserInfo, m *protojson.LeaveGameReq) {
	gameServer := g.GetGameServe(m.GameId)
	if gameServer == nil {
		global.G_LOG.Error("游戏不存在", zap.Uint64("gameId", m.GameId))
		return
	}
	gameServer.LeaveGameReq(userInfo, m)
}
