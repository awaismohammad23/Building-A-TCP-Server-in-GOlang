//building a custom tcp server in golang
package main

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddress string
	ln net.Listener
	QuitChanel chan struct{}
}

func NewServer(listenAddress string)*Server{

	return &Server{
		listenAddress: listenAddress,
		QuitChanel: make(chan struct{}),
	}
}

//function to start the server
func(s *Server) Start() error{
	ln,err := net.Listen("tcp",s.listenAddress)

	if err != nil{
		return err
	}

	defer ln.Close()
	s.ln = ln

	<- s.QuitChanel

	return nil	
}

//accept loop which will accept the connection and handle the connection
func(s *Server) acceptloop(){
	for{
		conn,err := s.ln.Accept()
		if err != nil{
			fmt.Println("accept error",err)
			continue
		}
		go s.readloop(conn)
	}
}

func (s*Server) readloop(conn net.Conn){

	defer conn.Close()
	buf := make([]byte,2048)

	for{
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("read error",err)
			continue
		}

		msg := buf[:n]
		fmt.Println("Received message: ",string(msg))
	}
}

func main()
{
	Server := NewServer(":3000")
	Server.Start()
}