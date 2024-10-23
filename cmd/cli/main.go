package main

import (
	"fmt"

	"github.com/feature-MaybeCoder/go-brocker/internal/messages"
)

func main() {
	reader := messages.JsonFileMessagesReader{}
	messages := reader.ReadMessagesGroup()
	fmt.Println(messages)
}
