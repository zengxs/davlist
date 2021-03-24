//go:generate yarn
//go:generate yarn build:fe:prod

// 不要将 package 命名为 vc
// 会与 vercel builder 产生冲突
package handler

import (
	"net/http"

	"davlist/internal/app"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app.DavlistMux.ServeHTTP(w, r)
}
