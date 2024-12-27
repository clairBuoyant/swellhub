package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed dist
var Dist embed.FS

func SPAHandler() http.HandlerFunc {
	staticFS, err := fs.Sub(Dist, "dist")
	if err != nil {
		panic(fmt.Errorf("failed getting site files: %w", err))
	}
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFS.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		if err == nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		http.FileServer(http.FS(staticFS)).ServeHTTP(w, r)
	}
}
