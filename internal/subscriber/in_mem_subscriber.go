package subscriber

import (
	main_config "github.com/feature-MaybeCoder/go-brocker/internal/config"
	"github.com/feature-MaybeCoder/go-brocker/internal/queue"
	"github.com/feature-MaybeCoder/go-brocker/internal/sender"
)

type inMemSubscriber struct {
	queues          []queue.Queue
	messages_sender sender.Sender
}

func (ims *inMemSubscriber) RunSendingLoop() {
	current_queue_index := 0
	for {
		err := 0
		current_queue := ims.queues[current_queue_index]
		current_attempt := 0
		message := current_queue.Pop()
		for current_attempt < main_config.MainConfig.MaxSendMessageRetries {
			err := ims.messages_sender.SendMessage(message)
			if err != nil {
				current_attempt++
				continue
			}
			break
		}
		if err != 0 {
			err := current_queue.PushMessage(message)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}

func NewInMemSubscriber(queues []queue.Queue, message_sender sender.Sender) Subscriber {
	return &inMemSubscriber{queues, message_sender}
}
