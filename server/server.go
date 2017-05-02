package server

import "net/http"

func Start(addr string) {
	http.HandleFunc("/stockquote", handler)
	http.ListenAndServe(addr, nil)
}
