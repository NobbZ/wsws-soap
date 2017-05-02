package server

import "net/http"

// Start starts a server listening on `addr` for incomming requests.
func Start(addr string) {
	http.HandleFunc("/stockquote", handler)
	http.ListenAndServe(addr, nil)
}
