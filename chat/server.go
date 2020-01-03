package main

import (
	"fmt"
	"net"
	"time"
)

//定义一个全局的channel，用于处理从各个客户端读到的消息
var message = make(chan []byte)

//定义一个结构体userInfo，用于存储每位聊天室用户的信息（名称+用户各自的管道C）
type userInfo struct {
	name   string
	C      chan []byte
	notice chan []byte //用于广播用户进入或退出当前聊天室的信息
}

//定义一个map，用于存储聊天室中所有在线的用户和用户信息
var onlineUsers = make(map[string]userInfo)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8011")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	fmt.Println("够浪聊天室-服务器已启动")

	fmt.Println("正在监听客户端连接请求……")

	//启动管家go程，不断监听全局channel————message
	go manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}
		fmt.Printf("地址为[%v]的客户端已连接成功\n", conn.RemoteAddr())
		// 如果监听到连接请求并成功以后，
		// 服务器进入下面的go程，
		// 在该go程中处理服务器和该客户端之间的读写或其他事件
		// 与此同时，服务器在主go程中回去继续监听着其他客户端的连接请求
		go HandleConnect(conn)
	}

}

// 这个函数完成服务器对一个客户端的整套处理流程
func HandleConnect(conn net.Conn) {
	defer conn.Close()
	// 管道overTime用于处理超时
	overTime := make(chan bool)

	// 用于存储用户名信息
	buf1 := make([]byte, 4096)
	n, err := conn.Read(buf1)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}
	userName := string(buf1[:n]) //n-1是为了去掉末尾的\n
	perC := make(chan []byte)
	perNewUser := make(chan []byte)
	user := userInfo{name: userName, C: perC, notice: perNewUser}
	onlineUsers[conn.RemoteAddr().String()] = user
	fmt.Printf("用户[%s]注册成功\n", userName)
	_, _ = conn.Write([]byte("💟💓💖💞💛你好," + userName + ",欢迎来到『够浪』™聊天室,请畅所欲言！💝💘💗💕💗"))
	//广播通知。遍历map
	go func() {
		for _, v := range onlineUsers {
			v.notice <- []byte("🤵用户[" + userName + "]已加入当前聊天室\n")
		}
	}()

	//监听每位用户自己的channel
	go func() {
		for {
			select {
			case msg1 := <-user.notice:
				_, _ = conn.Write(msg1)
			case msg2 := <-user.C:
				_, _ = conn.Write(msg2)

			}
		}
	}()

	//循环读取客户端发来的消息
	go func() {
		buf2 := make([]byte, 4096)
		for {
			n, err := conn.Read(buf2)
			//用于存储当前与服务器通信的客户端上的那个同户名
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			switch {
			case n == 0:
				fmt.Println(conn.RemoteAddr(), "已断开连接")
				for _, v := range onlineUsers {
					if thisUser != "" {
						v.notice <- []byte("💨用户[" + thisUser + "]已退出当前聊天室\n")
					}

				}
				delete(onlineUsers, conn.RemoteAddr().String())
				return
			case string(buf2[:n]) == "who\n":
				_, _ = conn.Write([]byte("当前在线用户:\n"))
				for _, v := range onlineUsers {
					//fmt.Println(v.name)
					_, _ = conn.Write([]byte("🟢" + v.name + "\n"))
				}
			case len(string(buf2[:n])) > 7 && string(buf2[:n])[:7] == "rename|":
				//n-1去掉buf2里的空格
				onlineUsers[conn.RemoteAddr().String()] = userInfo{name: string(buf2[:n-1])[7:], C: perC, notice: perNewUser}
				_, _ = conn.Write([]byte("您已成功修改用户名！\n"))
			}

			if err != nil {
				fmt.Println("conn.Read error:", err)
				return
			}

			var msg []byte
			if buf2[0] != 10 && string(buf2[:n]) != "who\n" {
				if len(string(buf2[:n])) <= 7 || string(buf2[:n])[:7] != "rename|" {
					msg = append([]byte("💬["+thisUser+"]对大家说:"), buf2[:n]...)
				}

			} else {
				msg = nil
			}
			//
			overTime <- true
			message <- msg
		}

	}()

	for {
		select {
		case <-overTime:
		case <-time.After(time.Second * 60):
			_, _ = conn.Write([]byte("抱歉，由于长时间未发送聊天内容，您已被系统踢出"))
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			for _, v := range onlineUsers {
				if thisUser != "" {
					v.notice <- []byte("💨用户[" + thisUser + "]由于长时间未发送消息已被踢出当前聊天室\n")
				}
			}
			delete(onlineUsers, conn.RemoteAddr().String())
			return
		}
	}

}

//管家循环监听管道message
func manager() {
	for m := range message {
		for _, v := range onlineUsers {
			v.C <- m
		}
	}
}
