package json

// {
//     "hdr" : {
//         "subid" : 10200,
//         "act" : "user_pwd"
//     },
//     "body" : {
//         "dev" : "dev",
//         "version" : "version",
//         "descrip" : "descrip",
//         "systype" : "systype",
//         "type" : "type",
//         "username" : "username",
//         "checkinfo" : "checkinfo",
//         "expinfo" : "expinfo"
//     }
// }

// 新加坡那边的协议
type SGMesageHDR struct {
	SubID uint32 `json:"subid"`
	Act   string `json:"act"`
}
type SGMessage struct {
	Hdr  SGMesageHDR `json:"hdr"`
	Body interface{} `json:"body"`
}

// -----------------------------------------------------------
type UserInfo struct {
	Nickanme string
	Uid      uint64
}

type MesageHeader struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginRsp struct {
	MesageHeader
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}

type RegisterReq struct {
	Account         string `json:"account"`
	Password        string `json:"password"`
	Nickname        string `json:"nickname"`
	ConfrimPassword string `json:"confrimPassword"`
}

type RegisterRsp struct {
	MesageHeader
}

type JoinGameReq struct {
	GameId uint64 `json:"gameId"`
}

type JoinGameRsp struct {
	MesageHeader
	// 对应的游戏信息
}

type LeaveGameReq struct {
	GameId uint64 `json:"gameId"`
}

type LeaveGameRsp struct {
	MesageHeader
}
