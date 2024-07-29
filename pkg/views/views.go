package views

import (
	"html/template"
	"net/http"
	"path/filepath"
	"webapp/pkg/logger"
)

func Index(w http.ResponseWriter, r *http.Request) {
    lp := filepath.Join("..", "pkg", "static", "templates", "index.html")
    tmpl, err := template.ParseFiles(lp)
    if err != nil {
        logger.Error.Printf("Error parsing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    data := struct {
        Owner string
    }{
        Owner: "Paul",
    }

    if err := tmpl.Execute(w, data); err != nil {
        logger.Error.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
