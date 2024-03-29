package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//儿子你好
//woxianggua...hahaaah
//hhhhhh

//IM功能：Instant Messaging即时通讯

//聊天系统（客户端）
func main() {
	Start(os.Args[1])
}
func Start(tcpAddrStr string) {
	//根据输入的ip加端口生成tcp地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpAddrStr)
	if err != nil {
		log.Printf("Resolve tcp add failed：%v\n", err)
		return
	}
	//向服务器拨号,成功返回一个conn
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("dial to server failed:%v\n", err)
		return
	}
	//向服务器发送消息
	go SendMsg(conn)
	//接收来自服务器端的广播消息
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed: %v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}
		fmt.Println(string(buf[0:length]))
	}
}

//向服务器端发消息
func SendMsg(conn net.Conn) {
	username := conn.LocalAddr().String()
	for {
		var input string
		//接收输⼊消息，放到input变量中
		fmt.Scanln(&input)
		if input == "/q" || input == "/quit" {
			fmt.Println("bye...")
			conn.Close()
			os.Exit(0)
		}
		//
		if len(input) > 0 {
			msg := username + " say:" + input
			_, err := conn.Write([]byte(msg))
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}
