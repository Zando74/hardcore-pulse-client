package secondary

import (
	"log"
	"os"
	"path/filepath"
)

type FileFinderImpl struct {}

func (f *FileFinderImpl) Find(folderPath string) ([]string, error) {
	var results []string

	err := filepath.WalkDir(folderPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			log.Println("[DEATHLOG-TRACKER] error on", path, ":", err)
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if d.Name() == "Deathlog.lua" {
			results = append(results, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}