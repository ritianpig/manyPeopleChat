package process

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes jsonMarshal err", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes jsonMarshal mes err", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes 发送失败 err", err.Error())
		return
	}
	return
}
