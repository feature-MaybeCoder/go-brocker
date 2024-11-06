package queue

import "github.com/feature-MaybeCoder/go-brocker/internal/models"

type Queue interface {
	PushMessage(models.Message) error
	Pop() models.Message
}

type QueuesManager interface {
	CreateQueue(name string) (Queue, error)
	GetQueue(name string) (Queue, bool)
	RecoverQueues() error
	persistQueue(Queue) error
}
