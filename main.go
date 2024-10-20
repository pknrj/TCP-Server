package main

import (
	"log"
	"net"
	"os"
	"time"
)

type TCPServer struct {
	add 	string 
	lis 	net.Listener
}

func NewTcpServer(add string) *TCPServer{
	return &TCPServer{
		add : add,
	}
}

func (s *TCPServer) Start() error {
	lis , err := net.Listen("tcp" , s.add) 
	log.Println("server started at port " , s.add)
	if err != nil {
		return err
	}
	defer lis.Close()
	s.lis = lis
	s.acceptConnection()
	return nil
}

func (s *TCPServer) acceptConnection(){
	for {
		c , err := s.lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.handleConnections(c)
	}
}

func (s *TCPServer) handleConnections(conn net.Conn){
	log.Println("request from remote server : " , conn.RemoteAddr().String())
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

	server := NewTcpServer(port)
	if err := server.Start() ; err != nil {
		log.Fatal(err)
	}

}