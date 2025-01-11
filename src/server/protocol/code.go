package protocol

const (
	FREAM_MSG = 10200
)

const (
	CODE_SUCCESS    = 200
	CODE_FAIL       = 500
	CODE_LOGIN_FAIL = 501
)

const (
	GAME_STATUS_NONE    = 0
	GAME_STATUS_LOADING = 1
	GAME_STATUS_RUNNING = 2
	// 维护
	GAME_STATUS_MAINTAIN = 3
	// 关闭
	GAME_STATUS_CLOSE = 4
)
