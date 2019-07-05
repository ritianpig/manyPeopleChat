package process

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"fmt"
	"io"
	"net"
)

// 定义服务器调度事件结构体,封装服务器调度相关方法
type Processor struct {
	Conn net.Conn
}

// 服务器事件调度
func (this *Processor) ServerProcess(mes *message.Message) (err error) {
	up := &UserProcess{
		Conn: this.Conn,
	}
	switch mes.Type {
	case message.LoginMesType:
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，暂时无法处理")
	}
	return
}

func (this *Processor) Process2() (err error) {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	// 循环读取客户端发送的数据包
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也即将退出...")
			} else {
				fmt.Println("readPkg()err =", err)
			}
			return err
		}

		err = this.ServerProcess(&mes)
		if err != nil {
			fmt.Println("process request serverProcess err=", err)
			return err
		}
	}
}
