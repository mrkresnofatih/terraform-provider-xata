package utilty

import (
	"fmt"
	"log"
)

func LogInfo(format string, v ...any) {
	LogBase("INFO ", format, v)
}

func LogError(format string, v ...any) {
	LogBase("ERROR", format, v)
}

func LogWarn(format string, v ...any) {
	LogBase("WARN ", format, v)
}

func LogBase(logLevel, format string, v ...any) {
	message := fmt.Sprintf(format, v)
	log.Printf("[%s] %s", logLevel, message)
}
