package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 定义结构体封装服务器对redis数据库的操作方法
type UserDao struct {
	Pool *redis.Pool
}

var (
	MyUserDao *UserDao
)

// NEwUserDao函数指定结构体操作哪个redis链接
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

// getUserById方法,服务器通过Id来查找指定的id是否存在用户信息
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.String(conn.Do("HGET", "user", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &message.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("getUserById Unmarshal err", err)
		return
	}
	return
}

// 校验用户登录信息
func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "user", user.UserId, string(data))
	if err != nil {
		fmt.Println("注册信息入库出错 err=", err)
		return
	}
	return
}