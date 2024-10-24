package main

import (
	"fmt"

	"github.com/feature-MaybeCoder/go-brocker/internal/reader"
)

func main() {
	reader := reader.JsonFileMessagesReader{}
	messages := reader.ReadMessagesGroup()
	fmt.Println(messages)
}
