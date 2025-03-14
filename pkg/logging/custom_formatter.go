package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func CustomLogFormatter(param gin.LogFormatterParams) string {
	_, file, line, ok := runtime.Caller(4) // Adjust depth if needed
	if !ok {
		file = "unknown"
		line = 0
	}

	// Format log message
	return fmt.Sprintf("[%s] %s | %3d | %13v | %-7s %s | %s:%d\n",
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		param.ClientIP,
		param.StatusCode,
		param.Latency,
		param.Method,
		param.Path,
		file, line,
	)
}
