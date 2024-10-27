package sender

import (
	"fmt"

	"github.com/feature-MaybeCoder/go-brocker/internal/models"
)

type dummySender struct {
}

func (ds *dummySender) SendMessage(message models.Message) error {
	fmt.Println(message)
	return nil
}

func NewDummySender() Sender {
	return &dummySender{}
}
