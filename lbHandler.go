package mybalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

func lbHandlerGenerate(cfg Config) func(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex
	var idx int = 0

	// lbHandler is the handler for loadbalancing
	lbHandler := func(w http.ResponseWriter, r *http.Request) {
		maxLen := len(cfg.Backends)

		// Round Robin
		mu.Lock()

		//getting current backend and if its dead we move to next idx
		//-This is an Active check for the backend - we do this for every request
		currentBackend := &cfg.Backends[idx%maxLen]
		if currentBackend.GetIsDead() {
			idx++
		}

		//choosing next backend server
		targetURL, err := url.Parse(cfg.Backends[idx%maxLen].URL)
		if err != nil {
			log.Fatal(err.Error())
		}
		idx++

		mu.Unlock()

		reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
		reverseProxy.ServeHTTP(w, r)
	}

	return lbHandler
}
