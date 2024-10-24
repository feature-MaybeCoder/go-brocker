package reader

import "github.com/feature-MaybeCoder/go-brocker/internal/models"

type MessagesReader interface {
	ReadMessagesGroup() models.MessagesGroup
}
