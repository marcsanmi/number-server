package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/marcsanmi/number-server/internal/network"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var (
	host string
	port string
)

func init() {
	// Initialize all the command line arguments
	const (
		defPort = "4000"
		defHost = "0.0.0.0"
	)
	flag.StringVar(&host, "host", defHost, "Host")
	flag.StringVar(&port, "port", defPort, "Port")
	flag.Parse()
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go printOutput(conn)
	parseInput(conn)
}

func parseInput(conn *net.TCPConn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Send message: ")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		err = network.Write(conn, text)
		if err != nil {
			log.Println(err)
		}
	}
}

func printOutput(conn *net.TCPConn) {
	for {
		message, err := network.Read(conn)
		// Receiving EOF means that the connection has been closed
		if err == io.EOF || strings.TrimSpace(string(message)) == "terminate" {
			// Close conn and exit
			fmt.Println("> Closing connection")
			conn.Close()
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(message)
	}
}
