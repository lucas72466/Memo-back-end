package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
)

type LogLevel int

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

const (
	DefaultIPFieldName          = "client_ip"
	DefaultUsernameFieldName    = "user_name"
	DefaultFileAndLineFieldName = "file:line"
)

var defaultLogger *logrus.Logger

func InitDefaultLogger() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	if mode := gin.Mode(); mode == gin.DebugMode {
		logger.SetLevel(logrus.TraceLevel)
		setDebugModeOutPut(logger)
	} else {
		logger.SetLevel(logrus.InfoLevel)
		// TODO release mode output
	}

	defaultLogger = logger
}

func LogWithContext(c *gin.Context, level LogLevel, msg interface{}, extraField map[string]interface{}) {
	extraField = injectDefaultContextInfo(c, extraField)

	defaultLogger.WithFields(extraField).Log(logrus.Level(level), msg)
}

func injectDefaultContextInfo(c *gin.Context, extraField map[string]interface{}) map[string]interface{} {
	if extraField == nil {
		extraField = make(map[string]interface{})
	}

	if gin.Mode() == gin.DebugMode {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			extraField[DefaultFileAndLineFieldName] = fmt.Sprintf("%s:%d", file, line)
		}
	}

	extraField[DefaultIPFieldName] = c.ClientIP()
	if userInfo := GetUserTokenInfoFromContextSilent(c); userInfo != nil {
		extraField[DefaultUsernameFieldName] = userInfo.UserName
	}

	return extraField
}

func setDebugModeOutPut(logger *logrus.Logger) {
	stdOut := os.Stdout
	localFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	logger.SetOutput(io.MultiWriter(stdOut, localFile))
}
