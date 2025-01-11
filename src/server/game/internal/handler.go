package internal

import (
	"reflect"
	"server/base"
	"server/global"
	protojson "server/protocol/json"

	"github.com/nlmayday/nlleaf/gate"
	"github.com/nlmayday/nlleaf/log"
	"go.uber.org/zap"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&protojson.SGMessage{}, handelAllMessage)

	handler(&protojson.JoinGameReq{}, handelJoinGame)
	handler(&protojson.LeaveGameReq{}, handelLeaveGame)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handelAllMessage(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*protojson.SGMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("m.Hdr.Act = %v, string(m.Body) = %v", m.Hdr.Act, (m.Body))

	// 给发送者回应一个 Hello 消息
	// var req protojson.SGMesageHDR
	// req.Act = "user_pwd"
	// req.SubID = 10201
	// var data protojson.SGMessage
	// data.Body = "testasdfasdfasdfasdf"
	// data.Hdr = req
	// a.WriteMsg(&data)
	_ = a
}

func handelJoinGame(args []interface{}) {
	m := args[0].(*protojson.JoinGameReq)
	// 消息的发送者
	a := args[1].(gate.Agent)

	userInfo := base.GetUserInfo(a)
	if userInfo == nil {
		global.G_LOG.Error("用户信息为空")
		// 返回错误 过期
		return
	}

	// 输出收到的消息的内容
	global.G_LOG.Debug("join", zap.Any("m", m))

	_ = a

	_game.JoinGameReq(userInfo, m)
}

func handelLeaveGame(args []interface{}) {
	m := args[0].(*protojson.LeaveGameReq)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	global.G_LOG.Debug("join", zap.Any("m", m))

	userInfo := base.GetUserInfo(a)
	if userInfo == nil {
		global.G_LOG.Error("用户信息为空")
		// 返回错误 过期
		return
	}
	_game.LeaveGameReq(userInfo, m)

	_ = a
}
