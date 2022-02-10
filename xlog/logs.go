/**
* @Author:Tristan
* @Date: 2021/11/11 3:17 下午
 */

package xlog

import (
	"context"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/shanlongpan/gin-mesh/config"
	"github.com/shanlongpan/gin-mesh/consts"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var Logger = logrus.New()

func init() {
	//把标准输入 discard
	Logger.Out = io.Discard

	//默认日志输出级别，可以设置为 error，写入另外一个文件 error 文件
	Logger.SetLevel(logrus.InfoLevel)

	//显示调用函数
	//Logger.SetReportCaller(true)
	statLogWriter, err := getLogWriter(config.Conf.LogName.StatLogFile)
	if err != nil {
		log.Fatalln(err)
	}

	errorLogWriter, err := getLogWriter(config.Conf.LogName.ErrorLogFile)
	if err != nil {
		log.Fatalln(err)
	}

	panicLogWriter, err := getLogWriter(config.Conf.LogName.PanicLogFile)
	if err != nil {
		log.Fatalln(err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  statLogWriter,
		logrus.FatalLevel: panicLogWriter,
		logrus.DebugLevel: statLogWriter,
		logrus.WarnLevel:  statLogWriter,
		logrus.ErrorLevel: errorLogWriter,
		logrus.PanicLevel: panicLogWriter,
	}, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger.AddHook(lfHook)
}
func createFile(fileName string) error {
	f, err := os.Create(fileName)
	defer f.Close()
	return err
}
func getLogWriter(logName string) (*rotatelogs.RotateLogs, error) {
	err := os.MkdirAll(config.Conf.LogDir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	fileName := path.Join(config.Conf.LogDir, logName+consts.Suffix)
	err = createFile(fileName)
	if err != nil {
		return nil, err
	}
	// 普通日志
	_, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	err = os.Chmod(fileName, 0777)
	if err != nil {
		return nil, err
	}
	// Set rotatelogs
	return rotatelogs.New(
		// Split file name
		path.Join(consts.LogFileDir, logName+"_%Y%m%d.log"),
		// Generate soft chain, point to the latest log file
		rotatelogs.WithLinkName(fileName),
		// Set maximum save time (7 days)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// Set log cutting interval (1 day)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
}

// 写入文件日志中间件
func LoggerToFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		ctx.Next()
		latencyTime := time.Since(startTime).String()
		// Request mode
		reqMethod := ctx.Request.Method
		// Request routing
		reqUri := ctx.Request.RequestURI
		// Status code
		statusCode := ctx.Writer.Status()
		// Request IP
		clientIP := ctx.ClientIP()
		traceId, _ := ctx.Get(consts.DefaultTraceIdHeader)

		// Log format
		Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"trace_id":     traceId,
		}).Info()
	}
}

func addTraceId(ctx context.Context) *logrus.Entry {
	traceId := ctx.Value(consts.DefaultTraceIdHeader)
	return Logger.WithFields(logrus.Fields{
		"trace_id": traceId,
	})
}

func Info(ctx context.Context, format string, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Info(format, args)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Infof(format, args)
}

func Infoln(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Infoln(args)
}

func Print(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Print(args)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(ctx context.Context, format string, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Printf(format, args)
}

func Println(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Println(args)
}

func Error(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Error(args)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Errorf(format, args)
}

func Errorln(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Errorln(args)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Fatal(args)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Fatalf(format, args)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Fatalln(args)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Panic(args)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(ctx context.Context, format string, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Panicf(format, args)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(ctx context.Context, args ...interface{}) {
	loggerEntry := addTraceId(ctx)
	loggerEntry.Println(args)
}

// Log to ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Logging to MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
