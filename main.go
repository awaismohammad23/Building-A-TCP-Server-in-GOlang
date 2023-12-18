//building a custom tcp server in golang
package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddress string
	ln            net.Listener
	QuitChanel    chan struct{}
	msgchn        chan []byte
}

func NewServer(listenAddress string) *Server {

	return &Server{
		listenAddress: listenAddress,
		QuitChanel:    make(chan struct{}),
		msgchn:        make(chan []byte, 10),
	}
}

//function to start the server
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)

	if err != nil {
		return err
	}

	defer ln.Close()
	s.ln = ln

	go s.acceptloop()

	<-s.QuitChanel

	close(s.msgchn)
	return nil
}

//accept loop which will accept the connection and handle the connection
func (s *Server) acceptloop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}

		fmt.Println("accepted connection from %s", conn.RemoteAddr())
		go s.readloop(conn)
	}
}

//function to read the message from the client
func (s *Server) readloop(conn net.Conn) {

	defer conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error", err)
			continue
		}

		s.msgchn <- buf[:n]
	}
}

func main() {
	Server := NewServer(":3000")

	go func() {

		for msg := range Server.msgchn {

			fmt.Println("Received message from connection: ", string(msg))
		}
	}()

	log.Fatal(Server.Start())
}
