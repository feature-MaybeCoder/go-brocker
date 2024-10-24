package queue

import (
	"github.com/feature-MaybeCoder/go-brocker/internal/models"
)

type inMemQueue struct {
	messages chan models.Message
}

func (imq *inMemQueue) PushMessage(message models.Message) error {
	imq.messages <- message
	return nil
}

func (imq *inMemQueue) Pop() models.Message {
	message := <-imq.messages
	return message
}

func NewInMemQueue(messages chan models.Message) inMemQueue {
	return inMemQueue{
		messages: messages,
	}
}
