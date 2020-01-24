package main

import (
	"flag"
	"github.com/marcsanmi/number-server/internal/generator"
	"github.com/marcsanmi/number-server/internal/server"
	"log"
	"os"
)

var (
	port              string
	concurrentClients int
)

func init() {
	// Initialize all the command line arguments
	const (
		defPort              = "4000"
		defConcurrentClients = 5
	)
	//flag.StringVar(&host, "host", defHost, "Host")
	flag.StringVar(&port, "port", defPort, "Port")
	flag.IntVar(&concurrentClients, "clients", defConcurrentClients, "Max concurrent clients")
	flag.Parse()
}

func main() {
	// Create the log file, path should be extracted
	file, err := os.Create("./tmp/numbers.log")
	//file, err := os.Create(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	msgChan := make(chan string)
	go generator.ProcessClientInput(file, msgChan)

	// Create the server
	s := server.NewServer()
	s.MessageChan = msgChan
	s.ConcurrentClients = concurrentClients
	err = s.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
	// Run the server
	s.Run()
}
