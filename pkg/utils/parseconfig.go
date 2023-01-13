package utils

import (
	"os"

	"github.com/IkehAkinyemi/zone-records-cli/pkg/model"
	"gopkg.in/yaml.v2"
)

// GetConfig read/parse YAML configuration file.
func GetConfig(fileDir string) (*model.Config, error) {
	file, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg model.Config

	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
