package server

import (
	"fmt"
	"github.com/marcsanmi/number-server/internal/commons"
	"github.com/marcsanmi/number-server/internal/network"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

const (
	ExitWord = "terminate"
)

type Server struct {
	connections       []net.Conn
	listener          net.Listener
	mutex             *sync.Mutex
	ConcurrentClients int
	MessageChan       chan string
}

func NewServer() *Server {
	return &Server{
		mutex: &sync.Mutex{},
	}
}

func (s *Server) Listen(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err == nil {
		s.listener = lis
	}
	log.Printf("> Listening on port %v", port)
	return err
}

func (s *Server) Run() {
	for {
		// Listen for connections
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Add connection and handle it limited by the maxConcurrentClients
		s.mutex.Lock()
		s.connections = append(s.connections, conn)
		s.mutex.Unlock()
		if len(s.connections) <= s.ConcurrentClients {
			go s.handleRequest(conn)
		} else {
			network.Write(conn, ExitWord)
			s.deleteConnection(conn)
			conn.Close()
		}
	}
}

func (s *Server) Close() {
	s.closeConnections()
	s.listener.Close()
}

func (s *Server) closeConnections() {
	for _, conn := range s.connections {
		conn.Close()
	}
}

// Handles incoming requests.
func (s *Server) handleRequest(conn net.Conn) {
	defer func() {
		s.deleteConnection(conn)
		conn.Close()
	}()
	for {
		message, err := network.Read(conn)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			break
		}
		message = strings.TrimSpace(string(message))
		fmt.Println("> Receiving:", message)
		// Terminate connection if receiving "terminate" or non-numeric input
		if message == ExitWord {
			s.mutex.Lock()
			s.broadcast(ExitWord)
			s.Close()
			s.mutex.Unlock()
			break
		}

		if !commons.IsNumeric(message) {
			network.Write(conn, ExitWord)
			break
		}

		// Parse message to be exactly of 9 digits length
		if len(message) > 9 {
			message = message[:9]
		} else {
			message = commons.LeftPad(message, "0", 9)
		}
		s.MessageChan <- message
	}
}

func (s *Server) deleteConnection(conn net.Conn) {
	var i int
	s.mutex.Lock()
	for i = range s.connections {
		if s.connections[i] == conn {
			break
		}
	}
	s.connections = append(s.connections[:i], s.connections[i+1:]...)
	s.mutex.Unlock()
}

func (s *Server) broadcast(msg string) {
	for i := range s.connections {
		err := network.Write(s.connections[i], msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Server) isValidInput(input string) bool {
	if !commons.IsNumeric(input) {
		return false
	}
	return true
}
