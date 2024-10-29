package main

import (
	"fmt"
	"sync"

	"github.com/feature-MaybeCoder/go-brocker/internal/models"
	"github.com/feature-MaybeCoder/go-brocker/internal/queue"
	"github.com/feature-MaybeCoder/go-brocker/internal/reader"
	"github.com/feature-MaybeCoder/go-brocker/internal/sender"
	"github.com/feature-MaybeCoder/go-brocker/internal/subscriber"
)

func main() {
	channel := make(chan models.Message, 10)
	dummy_sender := sender.NewDummySender()
	in_mem_queue := queue.NewInMemQueue(channel)
	in_mem_subscriber := subscriber.NewInMemSubscriber([]queue.Queue{&in_mem_queue}, dummy_sender)
	reader := reader.JsonFileMessagesReader{
		Queue: &in_mem_queue,
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go in_mem_subscriber.RunSendingLoop()
	go reader.RunReadingLoop()
	wg.Wait()
	fmt.Println()
}
