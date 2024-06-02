package app

import (
	"log"
	"os"
	"path/filepath"
)

// GetAppDir renvoie le chemin d’accès au répertoire utilisateur de l’application.
func GetAppDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("cannot find home directory: %v", err)
	}
	return filepath.Join(home, ".toggl-redmine")
}

// CreateAppDir crée le répertoire utilisateur de l’application s’il n’existe pas déjà.
func CreateAppDir() {
	d := GetAppDir()

	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err = os.Mkdir(d, 0o755); err != nil {
			log.Fatalf("cannot create app directory: %v", err)
		}
	}
}
