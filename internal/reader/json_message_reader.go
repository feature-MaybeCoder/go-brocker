package reader

import (
	"encoding/json"
	"os"

	main_config "github.com/feature-MaybeCoder/go-brocker/internal/config"
	"github.com/feature-MaybeCoder/go-brocker/internal/models"
	"github.com/feature-MaybeCoder/go-brocker/internal/queue"
)

type JsonFileMessagesReader struct {
	Queue queue.Queue
}

func (dmr *JsonFileMessagesReader) readMessage() models.Message {
	var messages_group models.MessagesGroup
	json_messages, err := os.ReadFile(
		main_config.MainConfig.FileInputDir.String(),
	)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(
		json_messages,
		&messages_group,
	)

	if err != nil {
		panic(err.Error())
	}

	return messages_group.Messages[0]
}

func (jmr *JsonFileMessagesReader) RunReadingLoop() {
	for {
		message := jmr.readMessage()
		jmr.Queue.PushMessage(message)

	}
}
