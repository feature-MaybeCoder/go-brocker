package main_config

import (
	"os"

	"github.com/chigopher/pathlib"
)

type mainConfig struct {
	BaseDir      pathlib.Path
	FileInputDir pathlib.Path
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

	return mainConfig{
		BaseDir:      *base_dir,
		FileInputDir: *file_input_dir,
	}
}

var MainConfig = parseConfig()
