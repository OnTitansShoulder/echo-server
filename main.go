package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"echo-server/handlers"
	"echo-server/processors"
	"echo-server/templates"
)

const (
	defaultPort = "5000"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Printf("$PORT is unset. Using default %s\n", defaultPort)
		port = defaultPort
	}

	// shared state
	allEchos := make(map[string]processors.Echo)
	echoChan := make(chan processors.Echo)
	defer close(echoChan)
	t := templates.ParseAllTemplates()

	// background routines
	go processors.TakeEchos(allEchos, echoChan)

	// api handlers
	http.DefaultClient.Timeout = time.Minute
	http.HandleFunc(handlers.EchoURLPath, handlers.EchoHandler(echoChan))
	http.HandleFunc(handlers.HealthURLPath, handlers.HealthCheckHandler())

	// view handlers
	http.HandleFunc(handlers.PastEchosViewURLPath, handlers.PastEchosViewHandler(allEchos, t))
	http.HandleFunc("/", handlers.PastEchosViewRedirectHandler())

	// start the server
	log.Printf("Listening on port %s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
