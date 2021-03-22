package app

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"davlist/frontend/public"
	"davlist/frontend/templates"
)

var DavlistMux *mux.Router

const templateFilename = "webpack-index.gohtml"

var tpl = template.Must(template.ParseFS(templates.Embedded, templateFilename))

func init() {
	DavlistMux = mux.NewRouter()

	DavlistMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tpl.Execute(w, nil)
	})

	fileServer := http.FileServer(http.FS(public.Embedded))
	DavlistMux.PathPrefix("/").Methods("GET").Handler(neuter(fileServer))
}

// 禁用 http.FileServer 的自动索引
func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
