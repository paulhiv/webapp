package routes

import (
	"path/filepath"
	"os"
	"net/http"
	"webapp/pkg/logger"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Log the 404 error in the same format as the request logs
	logMsg := logFmt(http.StatusNotFound, r.RemoteAddr, r.Method, r.URL.String(), "404 Not Found")
	logger.Error.Print(logMsg)
	//http.Error(w, "404 Not Found", http.StatusNotFound)
	http.ServeFile(w, r, filepath.Join("../pkg/static/templates", "404.html"))
}


func CustomFileServer(root http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crw := &CustomResponseWriter{ResponseWriter: w}
		_, err := root.Open(r.URL.Path)
		if os.IsNotExist(err) {
			NotFoundHandler(crw, r)
			return
		}
		http.FileServer(root).ServeHTTP(crw, r)
		if crw.StatusCode == 0 {
			crw.StatusCode = http.StatusOK
		}
		logMsg := logFmt(crw.StatusCode, r.RemoteAddr, "GET", r.URL.String(), "Static File")
		logger.Info.Print(logMsg)
	})
}
