package helper

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

var Validator = validator.New()

func Validate(i interface{}) error {
	return Validator.Struct(i)
}

// Response struct json
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Struct to give Empty Object
type EmptyObj struct {
}

// function Build Response
func BuildResponse(code string, message string, data interface{}) Response {
	res := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}

func UniqueCode() string {
	randomInt := math.MaxInt32
	rand := rand.Intn(randomInt)
	res := strconv.Itoa(rand)
	return res
}

func Log() {
	date := time.Now().Local().Format("2006-01-02")

	yesterday := time.Now().Local().AddDate(0, 0, -90).Format("2006-01-02")

	delErrLog := os.Remove("log/" + yesterday + ".log")
	delAllLog := os.Remove("log/" + yesterday + "-error.log")

	if delAllLog != nil && delErrLog != nil {
		fmt.Println(delAllLog.Error(), delErrLog.Error())
	}

	errorLog := handler.MustFileHandler("log/"+date+"-error.log", handler.WithLogLevels(slog.DangerLevels))

	allLog := handler.MustFileHandler("log/"+date+".log", handler.WithLogLevels(slog.AllLevels))

	// errorLog := handler.NewBuilder().
	// 	WithLogfile("log/error.log").
	// 	WithLogLevels(slog.DangerLevels).
	// 	WithRotateTime(rotatefile.EveryMinute).
	// 	Build()

	// allLog := handler.NewBuilder().
	// 	WithLogfile("log/log.log").
	// 	WithLogLevels(slog.AllLevels).
	// 	WithRotateTime(rotatefile.EveryMinute).
	// 	Build()

	// fc := rotatefile.NewFilesClear(func(c *rotatefile.CConfig) {
	// 	c.AddPattern("log/log.log")
	// 	c.BackupTime = 1
	// })

	// fc.DaemonClean(nil)
	// fc.StopDaemon()

	slog.PushHandler(errorLog)
	slog.PushHandler(allLog)
}

func Lpad(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}
