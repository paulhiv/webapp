package routes

import (
	"fmt"
	"net/http"
	"os"
	"webapp/pkg/logger"
)

const (
	reset    = "\033[0m"
	red      = "\033[31m"
	redbg    = "\033[41m"
	green    = "\033[32m"
	greenbg  = "\033[42m"
	orange   = "\033[38;5;208m"
	orangebg = "\033[48;5;208m"
	yellow   = "\033[33m"
	yellowbg = "\033[43m"
	blue     = "\033[34m"
	bluebg   = "\033[44m"
	cyan     = "\033[36m"
	cyanbg   = "\033[46m"
)

var methodColors = map[string]string{
	"GET":    bluebg,
	"POST":   cyanbg,
	"PUT":    yellowbg,
	"DELETE": redbg,
	"PATCH":  orangebg,
}

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (crw *CustomResponseWriter) WriteHeader(statusCode int) {
	if crw.StatusCode == 0 {
		crw.StatusCode = statusCode
	}
	crw.ResponseWriter.WriteHeader(statusCode)
}

func (crw *CustomResponseWriter) Write(b []byte) (int, error) {
	if crw.StatusCode == 0 {
		crw.StatusCode = http.StatusOK
	}
	return crw.ResponseWriter.Write(b)
}

func logFmt(status int, clientIp, method, url, pattern string) string {
	color, ok := methodColors[method]
	if !ok {
		color = reset
	}

	return fmt.Sprintf(
		"%s %d %s| %s |%s %s %s| %s | %s", color, status, reset, clientIp, color, method, reset, url, pattern,
	)
}

type CustomFileSystem struct {
	fs http.FileSystem
}

func (cfs CustomFileSystem) Open(name string) (http.File, error) {
	f, err := cfs.fs.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Error.Printf(" 404 File not found | %s", name)
		} else {
			logger.Error.Printf("Error opening file: %s, %v", name, err)
		}
		return nil, err
	}
	return f, nil
}
