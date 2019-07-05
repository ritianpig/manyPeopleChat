package process

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示二级菜单
func ShowMeu(userId int) {
	fmt.Printf("=====================恭喜%d登录成功=====================\n", userId)
	fmt.Println("\t\t\t 1　显示在线用户列表")
	fmt.Println("\t\t\t 2　发送消息")
	fmt.Println("\t\t\t 3　信息列表")
	fmt.Println("\t\t\t 4　退出系统")
	fmt.Println("\t\t\t 5　请选择(1-4):")

	var key int
	var content string
	fmt.Scanf("%d\n", &key)
	smsProcess := &SmsProcess{}

	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说的什么:)")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("请选择1-4")
	}

}

//和服务器保持不间断通讯
func KeepConn(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在和服务器保持通讯")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("保持通讯错误 err=", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("KeepConn Unmarshal err", err)
				continue
			}
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回未知的消息类型")
		}
	}
}
