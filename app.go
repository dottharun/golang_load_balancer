package mybalancer

import (
	"log"
	"net/http"
)

func Serve() {
	var cfg Config
	cfg.init()

	go healthCheck(cfg)

	server := http.Server{
		Addr:    ":" + cfg.Proxy.Port,
		Handler: http.HandlerFunc(lbHandlerGenerate(cfg)),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
