package handlers

import (
	"echo-server/processors"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	PastEchosViewURLPath = "/echos/view"
	echoTemplate         = "pastEchosView.gtpl.html"
)

type EchoListPageData struct {
	Echos []processors.Echo
}

func PastEchosViewRedirectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, PastEchosViewURLPath, 301)
	}
}

func PastEchosViewHandler(allEchos map[string]processors.Echo, t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, PastEchosViewURLPath) {
			http.NotFound(w, r)
			return
		}

		echos := make([]processors.Echo, 0, len(allEchos))
		for _, echo := range allEchos {
			echos = append(echos, echo)
		}
		data := EchoListPageData{
			Echos: echos,
		}

		err := t.ExecuteTemplate(w, echoTemplate, data)
		if err != nil {
			log.Printf("Error processing template=%s: %v\n", echoTemplate, err)
		}
	}
}
