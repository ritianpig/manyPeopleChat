package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

func outputOnlineUser() {
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
