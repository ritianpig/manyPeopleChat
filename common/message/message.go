package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMesType"
	RegisterResMesType      = "RegisterResMesType"
	NotifyUserStatusMesType = "NotifyUserStatusMesType"
	SmsMesType              = "SmsMesType"
)

const (
	UserOnline = iota
	UserOffline
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:userpwd`
	UserName string `json:username`
}

type LoginResMes struct {
	Code    int    `json:"code"`
	Error   string `json:error`
	UsersId []int  // 保存登录用户Id切片
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`
	User
}
