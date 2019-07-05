package process

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// 定义处理用户行为的结构体，用于对用户操作进行封装
type UserProcess struct {
}

// 用户登录处理方法
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// tcp 拨号链接
	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("josn.Marshal mes err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	// 发送数据包
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err=", err)
		return
	}

	// 读取服务器返回的数据包
	resMes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		fmt.Println("process Unmarshal err", err)
		return
	}
	if loginResMes.Code == 200 {
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		// fmt.Println("登录成功")
		fmt.Println("当前在线用户列表如下")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		go KeepConn(conn)
		for {
			ShowMeu(userId)
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

// 用户注册信息处理
func (this *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}
	defer conn.Close()

	// 定义发送的信息结构体
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 定义注册信息结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册信息发送错误 err=", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("注册信息接收服务器返回信息出错 err=", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功, 即将退出... 请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
