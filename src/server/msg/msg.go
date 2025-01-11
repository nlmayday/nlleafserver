package msg

import (
	"fmt"
	protojson "server/protocol/json"

	"github.com/nlmayday/nlleaf/network/json"
)

// var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	// 注册消息
	msgId := Processor.Register(&protojson.SGMessage{})
	fmt.Println(msgId)

	// 登陆
	Processor.Register(&protojson.LoginReq{})
	Processor.Register(&protojson.LoginRsp{})

	// 注册
	Processor.Register(&protojson.RegisterReq{})
	Processor.Register(&protojson.RegisterRsp{})

	// 加入游戏
	Processor.Register(&protojson.JoinGameReq{})
	Processor.Register(&protojson.JoinGameRsp{})

	// 退出游戏
	Processor.Register(&protojson.LeaveGameReq{})
	Processor.Register(&protojson.LeaveGameRsp{})

}
