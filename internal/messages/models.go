package messages

type Message struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	QueueName string `json:"queue_name"`
}

type MessagesGroup struct {
	Messages []Message `json:"messages"`
}
