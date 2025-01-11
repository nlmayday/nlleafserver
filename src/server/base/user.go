package base

import "github.com/nlmayday/nlleaf/gate"

type UserInfo struct {
	Uid uint64
}

func GetUserInfo(a gate.Agent) *UserInfo {
	if a.UserData() == nil {
		return nil
	}
	return a.UserData().(*UserInfo)
}

func SetUserInfo(a gate.Agent, user *UserInfo) {
	a.SetUserData(user)
}
