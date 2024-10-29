package sender

import "github.com/feature-MaybeCoder/go-brocker/internal/models"

type Sender interface {
	SendMessage(models.Message) error
}
