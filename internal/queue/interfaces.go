package queue

import "github.com/feature-MaybeCoder/go-brocker/internal/models"

type Queue interface {
	PushMessage(models.Message) error
	Pop() models.Message
}
