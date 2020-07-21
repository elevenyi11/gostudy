package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

//golang网络socket粘包问题的解决方法
//什么是粘包：百度上比较通俗的说法是指TCP协议中，发送方发送的若干包数据到接收方接收时粘成一包，从接收缓冲区看
//后一包数据的头紧接着前一包数据的尾

//Server

func TestSocketServer() {
	// 监听端口
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Printf("Listen Error: %s\n", err)
		return
	}
	// 监听循环
	for {
		// 接受客户端链接
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Accept Error: %s\n", err)
			continue
		}
		// 处理客户端链接
		go handleSocketConnection(conn)
	}
}

func handleSocketConnection(conn net.Conn) {
	// 关闭链接
	defer conn.Close()
	// 客户端
	fmt.Printf("Client: %s\n", conn.RemoteAddr())
	// 消息缓冲
	msgbuf := bytes.NewBuffer(make([]byte, 0, 10240))
	// 数据缓冲
	databuf := make([]byte, 4096)
	// 消息长度
	length := 0
	// 消息长度uint32
	ulength := uint32(0)
	// 数据循环
	for {
		// 读取数据
		n, err := conn.Read(databuf)
		if err == io.EOF {
			fmt.Printf("Client exit: %s\n", conn.RemoteAddr())
		}
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
			return
		}
		fmt.Println(databuf[:n])
		// 数据添加到消息缓冲
		n, err = msgbuf.Write(databuf[:n])
		if err != nil {
			fmt.Printf("Buffer write error: %s\n", err)
			return
		}
		// 消息分割循环
		for {
			// 消息头
			if length == 0 && msgbuf.Len() >= 4 {
				binary.Read(msgbuf, binary.LittleEndian, &ulength)
				length = int(ulength)
				// 检查超长消息
				if length > 10240 {
					fmt.Printf("Message too length: %d\n", length)
					return
				}
			}
			// 消息体
			if length > 0 && msgbuf.Len() >= length {
				fmt.Printf("Client messge: %s\n", string(msgbuf.Next(length)))
				length = 0
			} else {
				break
			}
		}
	}
}

//Client
func TestSocketClient() {
	// 链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	// 客户端信息
	fmt.Printf("Client: %s\n", conn.LocalAddr())
	// 消息缓冲
	msgbuf := bytes.NewBuffer(make([]byte, 0, 1024))
	// 消息内容
	message := []byte("我是utf-8的消息")
	// 消息长度
	messageLen := uint32(len(message))
	// 消息总长度
	mlen := 4 + len(message)
	// 写入5条消息
	for i := 0; i < 10; i++ {
		binary.Write(msgbuf, binary.LittleEndian, messageLen)
		msgbuf.Write(message)
	}
	// 单包发送一条消息
	conn.Write(msgbuf.Next(mlen))
	time.Sleep(time.Second)
	// 单包发送三条消息
	conn.Write(msgbuf.Next(mlen * 3))
	time.Sleep(time.Second)
	// 发送不完整的消息头
	conn.Write(msgbuf.Next(2))
	time.Sleep(time.Second)
	// 发送消息剩下部分
	conn.Write(msgbuf.Next(mlen - 2))
	time.Sleep(time.Second)
	// 发送不完整的消息体
	conn.Write(msgbuf.Next(mlen - 6))
	time.Sleep(time.Second)
	// 发送消息剩下部分
	conn.Write(msgbuf.Next(6))
	time.Sleep(time.Second)
	// 多段发送
	conn.Write(msgbuf.Next(mlen + 2))
	time.Sleep(time.Second)
	conn.Write(msgbuf.Next(-2 + mlen - 8))
	time.Sleep(time.Second)
	conn.Write(msgbuf.Next(8 + 1))
	time.Sleep(time.Second)
	conn.Write(msgbuf.Next(-1 + mlen + mlen))
	time.Sleep(time.Second)
	// 关闭链接
	conn.Close()
}