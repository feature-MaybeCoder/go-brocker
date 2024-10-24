package reader

import "github.com/feature-MaybeCoder/go-brocker/internal/models"

type MessagesReader interface {
	readMessage() models.MessagesGroup
	RunReadingLoop()
}
