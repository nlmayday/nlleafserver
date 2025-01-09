package gate

import (
	"server/game"
	"server/msg"
	protojson "server/proto/json"
)

func init() {
	msg.Processor.SetRouter(&protojson.LoginReq{}, game.ChanRPC)
}
