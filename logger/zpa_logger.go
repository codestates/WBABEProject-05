package logger

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// ZLog App 내 사용될 전역 로거
var ZapLog Logger

// lg 실제적인 zap 로거
var lg *zap.Logger

type LogConfigure interface {
	GetSettingValues() (path, level string, size, backup, age int)
}
type zapLogger struct {
}

func LoadZapLogger(conf LogConfigure) Logger {
	ZapLog = NewLogger(conf)
	return ZapLog
}

// 로거 초기화 컨피그 파라메터
func NewLogger(lcfg LogConfigure) (zapLogger *zapLogger) {
	path, level, size, backup, age := lcfg.GetSettingValues()

	now := time.Now()
	lPath := fmt.Sprintf("%s_%s.log", path, now.Format("2006-01-02"))
	// 설정 옵션
	writeSyncer := getLogWriter(lPath, size, backup, age)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(level)); err != nil {
		panic("logger load, fail")
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	// lg 생성
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return nil
}

func (z *zapLogger) Debug(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	lg.Debug("debug", zap.String("-", b.String()))
}

// Info is a convenient alias for Root().Info
func (z *zapLogger) Info(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	lg.Info("info", zap.String("-", b.String()))
}

// Warn is a convenient alias for Root().Warn
func (z *zapLogger) Warn(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	lg.Warn("warn", zap.String("-", b.String()))
}

// Error is a convenient alias for Root().Error
func (z *zapLogger) Error(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	lg.Error("error", zap.String("-", b.String()))
}

func (z *zapLogger) newWrittenBuffer(ctx []interface{}) *bytes.Buffer {
	var b bytes.Buffer
	for _, str := range ctx {
		b.WriteString(str.(string))
		b.WriteString(" ")
	}
	return &b
}

// GinLogger applied logger in gin
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		lg.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// gin 리커버리 대체 설정
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// encoder 옵션 설정
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
