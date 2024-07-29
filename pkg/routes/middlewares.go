package routes

import (
	"net/http"
	"strings"
	"webapp/pkg/logger"
)

func staticFileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crw := &CustomResponseWriter{ResponseWriter: w}

		if strings.HasPrefix(r.URL.Path, "/static/templates") {
			if r.Header.Get("X-Internal-Request") == "true" {
				next.ServeHTTP(crw, r)
				logMsg := logFmt(crw.StatusCode, r.RemoteAddr, r.Method, r.URL.String(), "Static File")
				logger.Info.Print(logMsg)
				return
			}

			http.Error(w, "403 Forbidden", http.StatusForbidden)
			logMsg := logFmt(http.StatusForbidden, r.RemoteAddr, r.Method, r.URL.String(), "403 Forbidden")
			logger.Info.Print(logMsg)
			return
		}

		next.ServeHTTP(crw, r)
	})
}

func logMiddleware(mux *http.ServeMux) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			crw := &CustomResponseWriter{ResponseWriter: w}
			next.ServeHTTP(crw, r)

			if crw.StatusCode == 0 {
				crw.StatusCode = http.StatusOK
			}

			_, pattern := mux.Handler(r)
			if pattern == "/" {
				pattern = "root"
			}

			logMsg := logFmt(crw.StatusCode, r.RemoteAddr, r.Method, r.URL.String(), pattern)
			logger.Info.Print(logMsg)
		})
	}
}
