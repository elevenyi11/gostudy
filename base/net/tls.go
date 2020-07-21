package net

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
)

//生成私钥：
//​openssl genrsa -out key.pem 2048
//生成证书：
//​openssl req -new -x509 -key key.pem -out cert.pem -days 3650
//https：
//提示: 访问请勿忘记使用https开头,否则chrome会下载一个文件如下:
//dotcoo-air:tls dotcoo$ cat /Users/dotcoo/Downloads/hello | xxd
//0000000: 1503 0100 0202 0a                        .......
func HelloTslServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func TestHelloServer() {
	http.HandleFunc("/hello", HelloTslServer)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//TLS Server：
func TestTLSServer() {
	log.SetFlags(log.Lshortfile)
	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":8000", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

//TLS Client：
func TestTLSClient() {
	log.SetFlags(log.Lshortfile)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:8000", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	println(string(buf[:n]))
}