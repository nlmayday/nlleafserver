package gate

import (
	"server/game"
	"server/login"
	"server/msg"
	protojson "server/protocol/json"
)

func init() {
	msg.Processor.SetRouter(&protojson.SGMessage{}, game.ChanRPC)

	msg.Processor.SetRouter(&protojson.LoginReq{}, login.ChanRPC)
	msg.Processor.SetRouter(&protojson.RegisterReq{}, login.ChanRPC)

	msg.Processor.SetRouter(&protojson.JoinGameReq{}, game.ChanRPC)
	msg.Processor.SetRouter(&protojson.LeaveGameReq{}, game.ChanRPC)
}
