package middleware

import (
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, h http.HandlerFunc)
}

func Reroute(w http.ResponseWriter, r *http.Request, hf http.HandlerFunc) {
	var hand Handler

	hf(w, r)

	hand.ServeHTTP(w, r, hf)
}
