package sender

import (
	"github.com/feature-MaybeCoder/go-brocker/internal/models"
)

type dummySender struct {
}

func (ds *dummySender) SendMessage(message models.Message) error {
	return nil
}

func NewDummySender() Sender {
	return &dummySender{}
}
