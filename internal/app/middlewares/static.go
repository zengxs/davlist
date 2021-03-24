package middlewares

import (
	"net/http"
	"os"

	"davlist/frontend/public"
)

var fs = http.FS(public.Embedded)

const (
	msg404 = "404 page not found"
	msg500 = "500 Internal Server Error"
)

func toHTTPError(err error, w http.ResponseWriter) {
	if os.IsNotExist(err) {
		http.Error(w, msg404, http.StatusNotFound)
	} else {
		http.Error(w, msg500, http.StatusInternalServerError)
	}
}

func QueryStaticHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/" {
			qs := r.URL.Query()
			if qs.Get("type") == "static" {
				path := qs.Get("path")

				f, err := fs.Open(path)
				if err != nil {
					toHTTPError(err, w)
					return
				}
				defer f.Close()

				d, err := f.Stat()
				if err != nil {
					toHTTPError(err, w)
					return
				}

				if d.IsDir() {
					http.NotFound(w, r)
					return
				}

				w.Header().Set("Vary", "Accept-Encoding")
				w.Header().Set("Cache-Control", "public, max-age=86400")
				http.ServeContent(w, r, d.Name(), d.ModTime(), f)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
