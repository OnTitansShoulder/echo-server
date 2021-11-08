package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	HealthURLPath = "/health/"
)

func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, HealthURLPath) {
			http.NotFound(w, r)
			return
		}

		fmt.Fprint(w, "ok")
	}
}
