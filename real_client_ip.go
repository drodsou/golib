package lib

import (
	"net/http"
	"strings"
)

// RealClientIP gets real ip of client even behind a reverse proxy
func RealClientIP(r *http.Request) string {

	forwarded := r.Header.Get("x-forwarded-for")

	if forwarded == "" {
		return strings.Split(r.RemoteAddr, ":")[0]
	} else {
		return forwarded
	}

}
