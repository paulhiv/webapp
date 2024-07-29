package logger

import (
	"log"
	"os"
	"fmt"
	"strings"
	"time"
	"runtime"
)

var (
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
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

const (
	dateWidth          = 10
	timeWidth          = 8
	maxFileNameWidth   = 15
	maxLineNumberWidth = 4
	maxLevelWidth      = 5
)

type customWriter struct {
	level string
}

func (cw *customWriter) Write(p []byte) (n int, err error) {
	_, file, line, _ := runtime.Caller(3)
	fileParts := strings.Split(file, "/")
	shortFile := fileParts[len(fileParts)-1]

	paddedDate := fmt.Sprintf("%-*s", dateWidth, time.Now().Format("02/01/2006"))
	paddedTime := fmt.Sprintf("%-*s", timeWidth, time.Now().Format("15:04:05"))
	paddedFile := fmt.Sprintf("%-*s", maxFileNameWidth, shortFile)
	paddedLine := fmt.Sprintf("%*d", maxLineNumberWidth, line)
	paddedLevel := fmt.Sprintf("%-*s", maxLevelWidth, cw.level)

var fmtLog string

switch cw.level {
case "INFO":
	fmtLog = fmt.Sprintf("%s%s%s %s | %s%s%s | %s:%s |%s", green, paddedTime, reset, paddedDate, green, paddedLevel, reset, paddedFile, paddedLine, p)

case "DEBUG":
	fmtLog = fmt.Sprintf("%s%s%s %s | %s%s%s | %s:%s |%s", blue, paddedTime, reset, paddedDate, blue, paddedLevel, reset, paddedFile, paddedLine, p)

case "ERROR":
	fmtLog = fmt.Sprintf("%s%s%s %s | %s%s%s | %s:%s |%s", red, paddedTime, reset, paddedDate, red, paddedLevel, reset, paddedFile, paddedLine, p)
}

return os.Stdout.Write([]byte(fmtLog))
}

func init() {
	Info = log.New(&customWriter{level: "INFO"}, "", 0)
	Debug = log.New(&customWriter{level: "DEBUG"}, "", 0)
	Error = log.New(&customWriter{level: "ERROR"}, "", 0)
}

