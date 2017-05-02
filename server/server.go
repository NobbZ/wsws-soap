package server

import (
	"log"
	"net/http"
)

// Start starts a server listening on `addr` for incomming requests.
func Start(addr string) {
	http.HandleFunc("/stockquote", handler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println("There was an error starting the handler")
	}
}
