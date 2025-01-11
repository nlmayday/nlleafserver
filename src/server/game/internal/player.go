package internal

import (
	"server/model"

	"github.com/nlmayday/nlleaf/gate"
)

type Player struct {
	GameId uint64
	Uid    uint64
	Agent  gate.Agent // 用户连接信息
	User   *model.User
}

func (p *Player) GetGameId() uint64 {
	return p.GameId
}

func (p *Player) GetUid() uint64 {
	return p.Uid
}

func (p *Player) GetAgent() gate.Agent {
	return p.Agent
}

func (p *Player) WriteMsg(msg interface{}) {
	if p.Agent != nil {
		p.Agent.WriteMsg(msg)
	}
}

func (p *Player) SetUser(user *model.User) {
	p.User = user
}

func (p *Player) GetUser() *model.User {
	return p.User
}

func (p *Player) SetAgent(agent gate.Agent) {
	if p.Agent != nil && p.Agent.LocalAddr().String() == agent.LocalAddr().String() {
		return
	}
	if p.Agent != nil {
		p.Agent.Close()
	}
	p.Agent = agent
}

func (P *Player) Release() {
	if P.Agent != nil {
		P.Agent.Close()
	}
}

func (P *Player) Logout() {
	if P.Agent != nil {
		P.Agent.Close()
	}

	// 记录数据 退出时间 等
}
