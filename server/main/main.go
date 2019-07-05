package main

import (
	"chatroom/common/redisdb"
	"chatroom/server/model"
	"chatroom/server/process"
	"fmt"
	"net"
	"time"
)

// 服务器调度入口
func control(conn net.Conn) (err error) {
	defer conn.Close()
	processor := &process.Processor{
		Conn: conn,
	}
	err = processor.Process2()
	if err != nil {
		fmt.Println("调度process2 err=", err)
		return
	}
	return
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(redisdb.Pool)
}

func main() {
	// 初始化redis连接池
	redisdb.InitPool("localhost:6379", 16, 0, 300*time.Second)
	// 初始化redis操作接口
	initUserDao()
	fmt.Println("服务端在8090端口监听...")
	// 监听tcp服务
	listen, err := net.Listen("tcp", "0.0.0.0:8090")
	defer listen.Close()
	if err != nil {
		fmt.Println("Listen err=", err)
		return
	}

	for {
		fmt.Println("监听成功等待客户端来链接...")
		// 等待用户链接,如果没有用户链接那么会一直阻塞
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			return
		}
		// 启动协程处理调度器
		go control(conn)

	}

}
