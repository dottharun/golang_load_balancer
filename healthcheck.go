package mybalancer

import (
	"log"
	"net"
	"net/url"
	"time"
)

// checks if the backend is alive.
func isBackendAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute*1)
	if err != nil {
		log.Printf("Unreachable to %v, error: %v", url.Host, err.Error())
		return false
	}
	defer conn.Close()
	return true
}

// healthCheck is a function for healthchecking all the backends
func healthCheck(cfg Config) {
	t := time.NewTicker(time.Minute * 1)

	for {
		select {
		case <-t.C:
			for i := range cfg.Backends {
				pingURL, err := url.Parse(cfg.Backends[i].URL)
				if err != nil {
					log.Fatal(err.Error())
				}

				isAlive := isBackendAlive(pingURL)
				cfg.Backends[i].SetDead(!isAlive)

				msg := "ok"
				if !isAlive {
					msg = "dead"
				}

				log.Printf("%v checked %v by healthcheck", cfg.Backends[i].URL, msg)
			}
		}
	}

}
