package hades

// AntdResult Antd Pro table list api
type AntdResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total" example:"100"`
}

type CommonResponse struct {
	ErrCode int    `json:"err_code" example:"0"`
	ErrMsg  string `json:"err_msg" example:"OK"`
	Data    any    `json:"data"`
}

// AntdLoginResult .
type AntdLoginResult struct {
	Status string `json:"status" enums:"ok,not ok"`
	Token  string `json:"token"`
}

// AntdCurrentUserResult .
type AntdCurrentUserResult struct {
	Name   string `json:"name"`
	Access string `json:"access" enums:"admin,agent"`
	Avatar string `json:"avatar"`
}
