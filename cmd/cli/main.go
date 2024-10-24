package main

import (
	"fmt"
	"sync"

	"github.com/feature-MaybeCoder/go-brocker/internal/models"
	"github.com/feature-MaybeCoder/go-brocker/internal/queue"
	"github.com/feature-MaybeCoder/go-brocker/internal/reader"
)

func main() {
	channel := make(chan models.Message, 10)
	in_mem_queue := queue.NewInMemQueue(channel)
	reader := reader.JsonFileMessagesReader{
		Queue: &in_mem_queue,
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go reader.RunReadingLoop()
	wg.Wait()
	fmt.Println()
}
