package routes

import (
	"net/http"
	"webapp/pkg/views"
	"webapp/pkg/logger"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Serve static files with 404 handling and access control
	fs := CustomFileServer(CustomFileSystem{fs: http.Dir("../pkg/static")})
	mux.Handle("/static/", staticFileMiddleware(http.StripPrefix("/static/", fs)))
	logger.Debug.Printf("Loaded static web assets")

	// Handle root requests with middleware
	mux.Handle("/", logMiddleware(mux)(http.HandlerFunc(views.Index)))
}

type CustomServeMux struct {
	mux             *http.ServeMux
	notFoundHandler http.Handler
}

func (csm *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, pattern := csm.mux.Handler(r)
	if handler == nil || pattern == "" {
		csm.notFoundHandler.ServeHTTP(w, r)
	} else {
		handler.ServeHTTP(w, r)
	}
}

func NewCustomServeMux(mux *http.ServeMux, notFoundHandler http.Handler) *CustomServeMux {
	return &CustomServeMux{
		mux:             mux,
		notFoundHandler: notFoundHandler,
	}
}

