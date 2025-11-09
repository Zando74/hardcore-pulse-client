package secondary

import (
	"deathlog-tracker/domain/entity"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type JSONGamePathRepository struct {
	configPath string
}

func NewJSONGamePathRepository(appName string) (*JSONGamePathRepository, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(home, ".config", appName)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	return &JSONGamePathRepository{
		configPath: filepath.Join(configDir, "config.json"),
	}, nil
}

func (r *JSONGamePathRepository) Save(p entity.GamePath) error {
	file, err := os.Create(r.configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(p)
}

func (r *JSONGamePathRepository) Find() (*entity.GamePath, error) {
	file, err := os.Open(r.configPath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var gp entity.GamePath
	if err := json.NewDecoder(file).Decode(&gp); err != nil {
		return nil, err
	}

	return &gp, nil
}
