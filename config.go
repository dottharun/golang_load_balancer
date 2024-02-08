package mybalancer

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Proxy is a reverse proxy, and means load balancer.
// port of the reverse-proxy through which client requests from frontend
type Proxy struct {
	Port string `json:"port"`
}

// Backend is servers which load balancer is transferred.
type Backend struct {
	URL    string `json:"url"`
	IsDead bool
	mu     sync.RWMutex //???why do we need a mutex on all backends in config
}

// SetDead updates the value of IsDead in Backend.
func (backend *Backend) SetDead(b bool) {
	backend.mu.Lock()
	backend.IsDead = b
	backend.mu.Unlock()
}

// GetIsDead returns the value of IsDead in Backend.
func (backend *Backend) GetIsDead() bool {
	backend.mu.RLock()
	isAlive := backend.IsDead
	backend.mu.RUnlock()
	return isAlive
}

// Config is the configuration of the user
type Config struct {
	Proxy    Proxy     `json:"proxy"`
	Backends []Backend `json:"backends"`
}

// init for Config object
func (c *Config) init() {
	data, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(data, &c)
}
