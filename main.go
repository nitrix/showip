package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func removePort(ip string) string {
	return strings.Split(ip, ":")[0]
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := removePort(r.RemoteAddr)

		if xForwardedFor := r.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
			xForwardedFor = removePort(strings.TrimSpace(strings.Split(xForwardedFor, ",")[0]))

			if xForwardedFor != "" && net.ParseIP(xForwardedFor) != nil {
				remoteAddr = xForwardedFor
			}
		} else if xRealIp := r.Header.Get("X-Real-IP"); xRealIp != "" {
			xRealIp = removePort(xRealIp)

			if xRealIp != "" && net.ParseIP(xRealIp) != nil {
				remoteAddr = xRealIp
			}
		}

		fmt.Fprint(w, remoteAddr)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
