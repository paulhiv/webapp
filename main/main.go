package main

import (
	"net/http"
	"webapp/pkg/routes"
	"webapp/pkg/logger"
)

func main() {
	mux := http.NewServeMux()
	customMux := routes.NewCustomServeMux(mux, http.HandlerFunc(routes.NotFoundHandler))
	routes.RegisterRoutes(mux)
	if err := http.ListenAndServe(":8080", customMux); err != nil {
		logger.Error.Fatalf("Server failed to start: %v", err)
	}
}
