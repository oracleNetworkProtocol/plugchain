package docs

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed index.tpl
var Docs embed.FS

// Handler returns an http handler that servers OpenAPI console for an OpenAPI spec at specURL.
func Handler(title, specURL string) http.HandlerFunc {
	t, err := template.ParseFS(Docs, "index.tpl")
	log.Println(t, err)
	return func(w http.ResponseWriter, req *http.Request) {
		t.Execute(w, struct {
			Title string
			URL   string
		}{
			title,
			specURL,
		})
	}
}
