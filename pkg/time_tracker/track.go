package timetracker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// TimeStore represents a structure for storing time values per key.
type TimeStore struct {
	mu    sync.Mutex
	store map[string]time.Time
}

// NewTimeStore creates a new TimeStore instance.
func NewTimeStore() *TimeStore {
	return &TimeStore{
		store: make(map[string]time.Time),
	}
}

// Set stores a time value for a given key.
func (ts *TimeStore) Set(key string, value time.Time) {
	if key == "" {
		log.Println("Error: Key cannot be empty.")
		return
	}
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.store[key] = value
	log.Printf("Set time for key '%s'.\n", key)
}

// Get retrieves the time value associated with a given key.
func (ts *TimeStore) Get(key string, format string) (string, bool) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	value, found := ts.store[key]
	if found {
		return value.Format(format), true
	}
	return "", false
}

// Delete removes a key and its associated time value from the store.
func (ts *TimeStore) Delete(key string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	delete(ts.store, key)
	log.Printf("Deleted key '%s'.\n", key)
}

// SaveToFile saves the TimeStore data to a JSON file.
func (ts *TimeStore) SaveToFile(filename string) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(ts.store); err != nil {
		return fmt.Errorf("Error encoding data to file: %w", err)
	}
	log.Println("Data saved to file.")
	return nil
}

// LoadFromFile loads the TimeStore data from a JSON file.
func (ts *TimeStore) LoadFromFile(filename string) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ts.store); err != nil {
		return fmt.Errorf("Error decoding data from file: %w", err)
	}
	log.Println("Data loaded from file.")
	return nil
}

