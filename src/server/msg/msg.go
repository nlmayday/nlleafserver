package msg

import (
	protojson "server/proto/json"

	"github.com/nlmayday/nlleaf/network/json"
)

// var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	// 注册消息
	Processor.Register(protojson.LoginReq{})
}
