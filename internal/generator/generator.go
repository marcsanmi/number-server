package generator

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	messageSet = make(map[string]bool)
	reportData = make(map[string]int)
	mutex      = &sync.Mutex{}
)

func ProcessClientInput(file *os.File, messageChan <-chan string) {
	// I use a mutex to avoid race conditions. Imagine multiple goroutines
	// have access to messageSet at the same time. The results could be wrong
	reportData = map[string]int{
		"unique":     0,
		"received":   0,
		"duplicated": 0,
	}

	go func() {
		getStatisticsEvery(10*time.Second, printReport)
	}()

	for message := range messageChan {
		// Check if the message already in the log file
		mutex.Lock()
		if !messageExists(message) {
			addMessage(message)
			reportData["unique"] += 1
			reportData["received"] += 1
			file.Write([]byte(fmt.Sprintf("%s\n", message)))
		} else {
			reportData["duplicated"] += 1
		}
		mutex.Unlock()
	}
}

func printReport() {
	mutex.Lock()
	fmt.Printf("Received %d unique numbers, %d duplicated. Unique total: %d\n",
		reportData["unique"], reportData["duplicated"], reportData["received"])
	reportData["unique"] = 0
	reportData["duplicated"] = 0
	mutex.Unlock()
}

func getStatisticsEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func messageExists(message string) bool {
	_, ok := messageSet[message]
	return ok
}

func addMessage(message string) {
	messageSet[message] = true
}
