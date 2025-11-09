package secondary

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
)

type PlayerHashRepositoryImpl struct {
	dbPath string
	db     []string
}

func NewPlayerHashRepositoryImpl(appName string) (*PlayerHashRepositoryImpl, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(home, ".config", appName)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	db := make([]string, 0)

	if _, err := os.Stat(filepath.Join(configDir, "db.json")); os.IsNotExist(err) {
		file, err := os.Create(filepath.Join(configDir, "db.json"))
		if err != nil {
			return nil, err
		}
		defer file.Close()
	} else {
		file, err := os.Open(filepath.Join(configDir, "db.json"))
		if err != nil {
			return nil, err
		}
		
		if err := json.NewDecoder(file).Decode(&db); err != nil {
			file.Close()
			file, err = os.Create(filepath.Join(configDir, "db.json"))
			if err != nil {
				return nil, err
			}
		}
		defer file.Close()
	}


	return &PlayerHashRepositoryImpl{
		dbPath: filepath.Join(configDir, "db.json"),
		db: db,
	}, nil
}

func (p *PlayerHashRepositoryImpl) SaveAll(playerhashs []string) error {
	p.db = append(p.db, playerhashs...)
	file, err := os.Create(p.dbPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(p.db)
}

func (p *PlayerHashRepositoryImpl) Exist(playerhash string) bool {
	return slices.Contains(p.db, playerhash)
}