package app

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"davlist/frontend/templates"
	"davlist/internal/app/middlewares"
)

var DavlistMux *mux.Router

const templateFilename = "webpack-index.gohtml"

var tpl = template.Must(template.ParseFS(templates.Embedded, templateFilename))

func init() {

	DavlistMux = mux.NewRouter()
	DavlistMux.Use(middlewares.QueryStaticHandler)

	DavlistMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tpl.Execute(w, nil)
	})
}
