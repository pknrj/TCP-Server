package main

import (
	"log"
	"net"
	"os"
	"time"
)


func handleConnections(conn net.Conn){
	buf := make([]byte , 1024)
	conn.Read(buf)
	time.Sleep(time.Second * 2)
	resp := "HTTP/1.1 200 OK\r\n\r\n" + "Hello from TCP server : " + conn.LocalAddr().String() + "\r\n"
	conn.Write([]byte(resp))
	conn.Close()
}


func main(){
	args := os.Args
	if len(args) == 1 {
		log.Println("Please provide port number ")
	}

	port := ":" + args[1]
	listener , err := net.Listen("tcp" , port)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		c , err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnections(c)
	}

}