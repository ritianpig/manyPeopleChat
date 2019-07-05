package main

import (
	"chatroom/client/process"
	"fmt"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	var key int

	for {
		fmt.Println("=====================欢迎登陆多人聊天系统=====================")
		fmt.Println("\t\t\t 1　登录聊天室")
		fmt.Println("\t\t\t 2　注册用户")
		fmt.Println("\t\t\t 3　退出系统")
		fmt.Println("\t\t\t 4　请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("Login err", err)
			}
		case 2:
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("Register err", err)
			}
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("输入错误，请重新输入")
		}

	}
}
