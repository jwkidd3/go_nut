package main

import (
	"fmt"
	"time"
)

func main() {
	bchan := make(chan string, 4)
	// We can send 4 values without a corresponding receive
	bchan <- "1"
	bchan <- "2"
	bchan <- "3"
	bchan <- "go!"
	go func() {
		bchan <- "go again"
	}()
	time.Sleep(1 * time.Second)
	for i := 1; i <= 5; i++ {
		fmt.Println(<-bchan)
	}
}
