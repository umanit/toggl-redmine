package app

import (
	"log"
	"os"
	"path/filepath"
)

const appDirName = "toggl-redmine"

// GetAppDir renvoie le chemin d’accès au répertoire de configuration de l’application,
// selon les conventions de l’OS ($XDG_CONFIG_HOME ou ~/.config sur Linux, ~/Library/Application
// Support sur macOS, %AppData% sur Windows).
func GetAppDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("cannot find config directory: %v", err)
	}
	return filepath.Join(dir, appDirName)
}

// legacyAppDir renvoie l’ancien emplacement du répertoire de l’application
// ($HOME/.toggl-redmine), utilisé avant la bascule vers les conventions de l’OS.
func legacyAppDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("cannot find home directory: %v", err)
	}
	return filepath.Join(home, ".toggl-redmine")
}

// migrateLegacyAppDir déplace le fichier de configuration de l’ancien emplacement vers le
// nouveau, s’il existe et n’a pas déjà été migré.
func migrateLegacyAppDir(newDir string) {
	oldConfig := filepath.Join(legacyAppDir(), "config.json")
	newConfig := filepath.Join(newDir, "config.json")

	if oldConfig == newConfig {
		return
	}

	if _, err := os.Stat(oldConfig); os.IsNotExist(err) {
		return
	}

	if _, err := os.Stat(newConfig); err == nil {
		return
	}

	if err := os.MkdirAll(newDir, 0o755); err != nil {
		log.Printf("cannot prepare new app directory for migration: %v", err)
		return
	}

	if err := os.Rename(oldConfig, newConfig); err != nil {
		log.Printf("cannot migrate config file from %s to %s: %v", oldConfig, newConfig, err)
		return
	}

	log.Printf("migrated config file from %s to %s", oldConfig, newConfig)
}

// CreateAppDir crée le répertoire de l’application s’il n’existe pas déjà, en migrant au
// préalable les données depuis l’ancien emplacement si besoin.
func CreateAppDir() {
	d := GetAppDir()

	migrateLegacyAppDir(d)

	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err = os.MkdirAll(d, 0o755); err != nil {
			log.Fatalf("cannot create app directory: %v", err)
		}
	}
}
