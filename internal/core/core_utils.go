package core

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type cached_rules struct {
	mu    sync.RWMutex
	rules map[string][]byte
}

type cached_config struct {
	debug_mode bool
	test_mode  bool
	hot_reload bool
	port       int
}

type config_handler struct {
	mu      sync.RWMutex
	config  *cached_config
	path    string
	watcher *fsnotify.Watcher
}

func (cc *cached_config) ConfigInMemory() {

}

func SetupLogger(debug bool) error {

	logFile, err := os.Create(filepath.Join("..", "..", "logs", fmt.Sprintf("TBE_%s.log", time.Now().Format(time.DateOnly))))
	if err != nil {
		return fmt.Errorf("Failed while attempting to create log file: %w", err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logWriter := io.NewOffsetWriter(logFile, 0)
	log.SetOutput(logWriter)

	if debug {
		log.SetPrefix("DEBUG: ")
	}

	return nil

}
