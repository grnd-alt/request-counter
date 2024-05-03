package database

import (
	"log"
	"os"
	"path/filepath"
	"sort"
)

func ApplyMigrations() {
	db := getDB()
	paths, err := filepath.Glob("./database/migrations/*.sql")
	if err != nil {
		log.Fatal("MIGRATIONS FAILED")
	}

	sort.Strings(paths)

	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatal(err)
		}
	}

}
