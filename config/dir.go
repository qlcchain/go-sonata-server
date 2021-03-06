/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package config

import (
	"os"
	"os/user"
	"path/filepath"
	osRuntime "runtime"

	log "github.com/sirupsen/logrus"
)

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

func dataDir() string {
	home := homeDir()
	dataDir := "Gsonata"
	if home != "" {
		if osRuntime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "Application Support", dataDir)
		} else if osRuntime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", dataDir)
		} else {
			return filepath.Join(home, dataDir)
		}
	}
	return ""
}

func joinDir(dir string) string {
	d := dataDir()
	if d == "" {
		d = "/tmp"
	}
	p := filepath.Join(d, dir)
	if err := createDirIfNotExist(p); err != nil {
		log.Error(err)
	}
	return p
}

func LogDir() string {
	return joinDir("log")
}

func DBDir() string {
	return joinDir("db")
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		return err
	}
	return nil
}
