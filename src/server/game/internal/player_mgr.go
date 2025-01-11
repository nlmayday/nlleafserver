package internal

import (
	"server/model"

	"github.com/nlmayday/nlleaf/gate"
)

type PlayerMgr struct {
	players map[uint64]*Player
}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		players: make(map[uint64]*Player),
	}
}

func (pm *PlayerMgr) AddPlayer(user *model.User, a gate.Agent) error {
	if p, ok := pm.players[user.ID]; ok {
		p.SetAgent(a)
		p.SetUser(user)
	} else {
		player := new(Player)
		player.Uid = user.ID
		player.SetAgent(a)
		player.SetUser(user)
		pm.players[user.ID] = player

	}
	return nil
}

func (pm *PlayerMgr) RemovePlayer(uid uint64) {
	if p, ok := pm.players[uid]; ok {
		p.Release()
		delete(pm.players, uid)
	}
}

func (pm *PlayerMgr) GetPlayer(uid uint64) *Player {
	if p, ok := pm.players[uid]; ok {
		return p
	}
	return nil
}

func (pm *PlayerMgr) GetPlayersByGameId() map[uint64]*Player {
	return nil
}
