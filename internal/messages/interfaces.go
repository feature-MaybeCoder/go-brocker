package messages

type MessagesReader interface {
	ReadMessagesGroup() MessagesGroup
}
