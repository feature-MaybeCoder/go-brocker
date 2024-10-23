package messages

import (
	"encoding/json"
	"os"

	main_config "github.com/feature-MaybeCoder/go-brocker/internal/config"
)

type JsonFileMessagesReader struct {
}

func (dmr *JsonFileMessagesReader) ReadMessagesGroup() MessagesGroup {
	var messages_group MessagesGroup
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

	return messages_group
}
