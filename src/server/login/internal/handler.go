package internal

import (
	"fmt"
	"reflect"
	"server/model"
	"server/protocol"
	protojson "server/protocol/json"

	"github.com/nlmayday/nlleaf/gate"
	"github.com/nlmayday/nlleaf/log"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&protojson.LoginReq{}, handelLogin)
	handler(&protojson.RegisterReq{}, handelRegister)
}

func handelRegister(args []interface{}) {
	m := args[0].(*protojson.RegisterReq)
	a := args[1].(gate.Agent)
	log.Debug("m = %v", m)
	a.WriteMsg(&protojson.RegisterRsp{})
	var user model.User
	user.Account = m.Account
	user.Password = m.Password
	user.Nickname = m.Nickname
	err := model.AddUser(&user)

	var req protojson.JoinGameRsp
	req.Code = protocol.CODE_SUCCESS
	req.Msg = "注册成功"
	if err != nil {
		log.Debug("err = %v", err)
		req.Code = protocol.CODE_FAIL
		req.Msg = err.Error()
	}

	a.WriteMsg(req)
}

func handelLogin(args []interface{}) {
	m := args[0].(*protojson.RegisterReq)
	a := args[1].(gate.Agent)
	log.Debug("m = %v", m)

	// 判断是否已经登陆
	if a.UserData() != nil {
		log.Debug("已经登陆")
		return
	}

	var req protojson.LoginRsp
	req.Code = protocol.CODE_SUCCESS
	req.Msg = "注册成功"

	user, err := model.UserLogin(m.Account, m.Password)
	if err != nil {
		req.Code = protocol.CODE_LOGIN_FAIL
		req.Msg = err.Error()
		log.Debug("err = %v", err)
		a.WriteMsg(req)
		return
	}

	a.SetUserData(user)

	req.Token = fmt.Sprintf("%d", user.ID)
	req.UserInfo = protojson.UserInfo{
		Nickanme: user.Nickname,
		Uid:      uint64(user.ID),
	}
	a.WriteMsg(req)

	// 加入到全局管理
}
