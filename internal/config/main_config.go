package main_config

import (
	"os"

	"github.com/chigopher/pathlib"
)

type mainConfig struct {
	BaseDir               pathlib.Path
	FileInputDir          pathlib.Path
	QueuesConfigDir       pathlib.Path
	MaxSendMessageRetries int
}

func parseConfig() mainConfig {
	// Get BaseDir
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	current_dir := pathlib.NewPath(ex)
	base_dir := current_dir.Parent().Parent()

	// Get file input path
	file_input_dir := base_dir.Join("input/messages.json")

	queues_config_dir := base_dir.Join("state/queues")

	return mainConfig{
		BaseDir:               *base_dir,
		FileInputDir:          *file_input_dir,
		QueuesConfigDir:       *queues_config_dir,
		MaxSendMessageRetries: 10,
	}
}

var MainConfig = parseConfig()
