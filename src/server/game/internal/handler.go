package internal

import (
	"reflect"
	protojson "server/proto/json"

	"github.com/nlmayday/nlleaf/gate"
	"github.com/nlmayday/nlleaf/log"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&protojson.LoginReq{}, handleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*protojson.LoginReq)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("account %v", m.Account)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&protojson.LoginReq{
		Account: "client",
	})
}
