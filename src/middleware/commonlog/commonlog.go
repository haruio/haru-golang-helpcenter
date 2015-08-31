package commonlog

import (
	"bytes"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// Instances a Logger middleware that will write the logs to gin.DefaultWriter
// By default gin.DefaultWriter = os.Stdout
func Logger() gin.HandlerFunc {
	return NewWithWriter(gin.DefaultWriter)
}

// Instance a Logger middleware with the specified writter buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func NewWithWriter(out io.Writer) gin.HandlerFunc {
	pool := &sync.Pool{
		New: func() interface{} {
			buf := new(bytes.Buffer)
			return buf
		},
	}
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start).String()
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		methodColor := colorForMethod(method)
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		w := pool.Get().(*bytes.Buffer)
		w.Reset()
		w.WriteString(end.Format("2006/01/02 - 15:04:05"))
		w.WriteString(" |")
		w.WriteString(statusColor)
		w.WriteString(strconv.Itoa(statusCode))
		w.WriteString(reset)
		w.WriteString("| ")
		w.WriteString(clientIP)
		w.WriteString("\t ")
		w.WriteString(latency)
		w.WriteString(" \t")
		w.WriteString(methodColor)
		w.WriteString(c.Request.Method)
		w.WriteString(reset)
		w.WriteString(comment)
		w.WriteString(" ")
		w.WriteString(c.Request.Proto)
		w.WriteString(" ")
		w.WriteString(path)
		w.WriteString(" ")
		w.WriteString("\t Size ")
		w.WriteString(strconv.Itoa(c.Writer.Size()))
		w.WriteString("\n")
		w.WriteTo(out)

		pool.Put(w)
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
