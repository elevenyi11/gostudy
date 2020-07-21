package rpc

import (

	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"time"
)

type Args struct {
	A,B int
}
type Arith int
func (this *Arith) Multiply(args *Args, reply *([]string)) error{
	*reply = append(*reply,"test:")
	*reply = append(*reply,strconv.Itoa(args.A))
	*reply = append(*reply,"*")
	*reply = append(*reply,strconv.Itoa(args.B))
	*reply = append(*reply,"=")
	*reply = append(*reply,strconv.Itoa(args.A * args.B))
	return nil
}

func TestRPC(){
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l,e := net.Listen("tcp",":1234")
	if e != nil{
		log.Fatal("listen error:",e)
	}
	go http.Serve(l,nil)

	time.Sleep(5*time.Second)
	client,err:= rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err != nil{
		log.Fatal("dialing:",err)
	}
	args := &Args{7,8}
	replay := make([]string,10)
	err = client.Call("Arith.Multiply",args,&replay)
	if err != nil{
		log.Fatal("call arith error:",err)
	}
	log.Println(replay)
}

func TestRPC2(){
	newServer := rpc.NewServer()
	newServer.Register(new(Arith))
	l, e:= net.Listen("tcp","127.0.0.1:1234")
	if e != nil{
		log.Fatalf("net.listen tcp:0:%v",e)
	}
	go newServer.Accept(l)
	newServer.HandleHTTP("/foo","/bar")
	time.Sleep(1*time.Second)
	address,err := net.ResolveTCPAddr("tcp","127.0.0.1:1234")
	if err != nil{
		panic(err)
	}
	conn,_:= net.DialTCP("tcp",nil,address)
	defer conn.Close()
	client := rpc.NewClient(conn)
	defer client.Close()
	args := &Args{7,8}
	replay := make([]string,10)
	err = client.Call("Arith.Multiply",args,&replay)
	if err != nil{
		log.Fatal("arith error:", err)
	}
	log.Println(replay)
}