//building a custom tcp server in golang
package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	From    string
	payLoad []byte
}

type Server struct {
	listenAddress string
	ln            net.Listener
	QuitChanel    chan struct{}
	msgchn        chan Message
}

func NewServer(listenAddress string) *Server {

	return &Server{
		listenAddress: listenAddress,
		QuitChanel:    make(chan struct{}),
		msgchn:        make(chan Message, 10),
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

		s.msgchn <- Message{
			From:    conn.RemoteAddr().String(),
			payLoad: buf[:n],
		}

		conn.Write([]byte("Message received :)\n"))
	}
}

func main() {
	Server := NewServer(":3000")

	go func() {

		for msg := range Server.msgchn {

			fmt.Printf("Received message from connection: (%s):%s\n", msg.From, string(msg.payLoad))
		}
	}()

	log.Fatal(Server.Start())
}
