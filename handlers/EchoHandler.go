package handlers

import (
	"echo-server/processors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	EchoURLPath = "/echo/"
)

var forwardedRemoteAddrHeaders = map[string]bool{
	"X-Forwarded-For": true, // heroku router
}

func EchoHandler(echoChan chan processors.Echo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, EchoURLPath) {
			http.NotFound(w, r)
			return
		}

		// extract the IP address only
		rAddr := r.RemoteAddr
		if len(rAddr) == 0 {
			log.Println("Warn: received request with empty remote address")
			return
		}
		IPAddr := strings.Split(rAddr, ":")[0]

		// check for IP override from router-attached Headers
		for name, values := range r.Header {
			if _, ok := forwardedRemoteAddrHeaders[name]; ok {
				if len(values) == 0 {
					log.Printf("Warn: header %s is present but has empty value\n", name)
					continue
				}
				IPAddr = values[0]
			}
		}

		now := time.Now()
		echo := processors.Echo{
			IP:          IPAddr,
			LastHitTime: now.Format(time.RubyDate),
			Count:       1,
		}
		echo.FirstHitTime = echo.LastHitTime
		echo.LastHitTimestamp = now.UnixMilli()
		echoChan <- echo

		fmt.Fprint(w, "ok")
	}
}
