package watcher

import (
	"context"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/mathieudottir/poker_tracker/backend/internal/services"
)

type FileWatcher struct {
	watcher       *fsnotify.Watcher
	importService *services.ImportService
	userID        int
	directory     string
	mu            sync.Mutex
	running       bool
	processedFiles map[string]bool
}

func NewFileWatcher(importService *services.ImportService, userID int, directory string) (*FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &FileWatcher{
		watcher:       watcher,
		importService: importService,
		userID:        userID,
		directory:     directory,
		processedFiles: make(map[string]bool),
	}, nil
}

func (fw *FileWatcher) Start(ctx context.Context) error {
	if err := fw.watcher.Add(fw.directory); err != nil {
		return err
	}

	fw.mu.Lock()
	fw.running = true
	fw.mu.Unlock()

	log.Printf("File watcher started for directory: %s", fw.directory)

	go fw.watch(ctx)

	return nil
}

func (fw *FileWatcher) watch(ctx context.Context) {
	debounceTimer := time.NewTimer(0)
	<-debounceTimer.C // Drain the initial timer

	pendingFiles := make(map[string]time.Time)

	for {
		select {
		case <-ctx.Done():
			fw.Stop()
			return

		case event, ok := <-fw.watcher.Events:
			if !ok {
				return
			}

			// Only process creation and write events for summary files
			if event.Op&(fsnotify.Create|fsnotify.Write) != 0 {
				if strings.HasSuffix(event.Name, "_summary.txt") {
					fw.mu.Lock()
					if !fw.processedFiles[event.Name] {
						pendingFiles[event.Name] = time.Now()
					}
					fw.mu.Unlock()

					// Reset debounce timer
					debounceTimer.Reset(2 * time.Second)
				}
			}

		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("File watcher error: %v", err)

		case <-debounceTimer.C:
			// Process pending files
			fw.mu.Lock()
			for summaryFile := range pendingFiles {
				if !fw.processedFiles[summaryFile] {
					go fw.processFile(ctx, summaryFile)
					fw.processedFiles[summaryFile] = true
				}
			}
			pendingFiles = make(map[string]time.Time)
			fw.mu.Unlock()
		}
	}
}

func (fw *FileWatcher) processFile(ctx context.Context, summaryFile string) {
	// Wait a bit to ensure file is fully written
	time.Sleep(500 * time.Millisecond)

	hhFile := strings.Replace(summaryFile, "_summary.txt", ".txt", 1)

	log.Printf("Processing new tournament: %s", summaryFile)

	if err := fw.importService.ImportTournament(ctx, fw.userID, hhFile, summaryFile); err != nil {
		log.Printf("Error importing tournament from %s: %v", summaryFile, err)
	} else {
		log.Printf("Successfully imported tournament from %s", summaryFile)
	}
}

func (fw *FileWatcher) Stop() {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	if fw.running {
		fw.watcher.Close()
		fw.running = false
		log.Printf("File watcher stopped")
	}
}

func (fw *FileWatcher) IsRunning() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	return fw.running
}

func (fw *FileWatcher) GetDirectory() string {
	return fw.directory
}

// WatcherManager manages multiple file watchers for different users
type WatcherManager struct {
	watchers map[int]*FileWatcher
	mu       sync.RWMutex
}

func NewWatcherManager() *WatcherManager {
	return &WatcherManager{
		watchers: make(map[int]*FileWatcher),
	}
}

func (wm *WatcherManager) StartWatcher(ctx context.Context, userID int, directory string, importService *services.ImportService) error {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	// Stop existing watcher if any
	if existing, ok := wm.watchers[userID]; ok {
		existing.Stop()
	}

	// Create and start new watcher
	watcher, err := NewFileWatcher(importService, userID, directory)
	if err != nil {
		return err
	}

	if err := watcher.Start(ctx); err != nil {
		return err
	}

	wm.watchers[userID] = watcher
	return nil
}

func (wm *WatcherManager) StopWatcher(userID int) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	if watcher, ok := wm.watchers[userID]; ok {
		watcher.Stop()
		delete(wm.watchers, userID)
	}
}

func (wm *WatcherManager) GetWatcher(userID int) *FileWatcher {
	wm.mu.RLock()
	defer wm.mu.RUnlock()
	return wm.watchers[userID]
}
