package global

import (
	"fmt"
	"gopkg.in/clog.v1"
	"io"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func InitLog(names []string) {
	if len(names) == 0 {
		panic("没有传入日志文件名称")
	}
	l := Cfg.Section("log")
	levelNames := map[string]clog.LEVEL{
		"trace": clog.TRACE,
		"info":  clog.INFO,
		"warn":  clog.WARN,
		"error": clog.ERROR,
		"fatal": clog.FATAL,
	}
	ll := strings.ToLower(strings.TrimSpace(l.Key("LEVEL").MustString("Trace")))
	level, ok := levelNames[ll]
	if !ok {
		log.Fatal(2, "日志级别错误: %v", ll)
	}
	frc := clog.FileRotationConfig{
		Rotate:   l.Key("ROTATE").MustBool(true),
		Daily:    l.Key("DAILY").MustBool(true),
		MaxSize:  1 << uint(l.Key("MAX_SIZE_SHIFT").MustInt(28)),
		MaxLines: l.Key("MAX_LINES").MustInt64(10000),
		MaxDays:  l.Key("MAX_DAYS").MustInt64(30),
	}
	Logs = make(map[string]*Logger, len(names))
	for _, name := range names {
		lr, err := clog.NewFileWriter(path.Join(path.Join(AppPath, "log"), name+".log"), frc)
		if err != nil {
			panic("log NewFileWriter err: " + err.Error())
		}
		Logs[name] = NewLogger3(lr, "", 0, level)
	}
}

type Logger struct {
	TRACE *log.Logger
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	FATAL *log.Logger
	level clog.LEVEL
}

func NewLogger(out io.Writer) *Logger {
	return NewLogger2(out, "", log.Ldate|log.Ltime)
}

func NewLogger2(out io.Writer, prefix string, flag int) *Logger {
	return NewLogger3(out, prefix, flag, clog.TRACE)
}

func NewLogger3(out io.Writer, prefix string, flag int, l clog.LEVEL) *Logger {
	return &Logger{
		TRACE: log.New(out, prefix, flag),
		INFO:  log.New(out, prefix, flag),
		WARN:  log.New(out, prefix, flag),
		ERROR: log.New(out, prefix, flag),
		FATAL: log.New(out, prefix, flag),
		level: l,
	}
}

func (l *Logger) Trace(format string, v ...interface{}) {
	if l.level <= clog.TRACE {
		l.TRACE.Output(2, write(clog.TRACE, 2, format, v...))
	}
	return
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.level <= clog.INFO {
		l.INFO.Output(2, write(clog.INFO, 2, format, v...))
	}
	return
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l.level <= clog.WARN {
		l.WARN.Output(2, write(clog.WARN, 2, format, v...))
	}
	return
}

func (l *Logger) Error(skip int, format string, v ...interface{}) {
	if l.level <= clog.ERROR {
		l.ERROR.Output(skip, write(clog.ERROR, skip, format, v...))
	}
	return
}

func (l *Logger) Fatal(skip int, format string, v ...interface{}) {
	if l.level <= clog.FATAL {
		l.FATAL.Output(skip, write(clog.FATAL, skip, format, v...))
	}
	return
}

var formats = map[clog.LEVEL]string{
	clog.TRACE: "[TRACE] ",
	clog.INFO:  "[ INFO] ",
	clog.WARN:  "[ WARN] ",
	clog.ERROR: "[ERROR] ",
	clog.FATAL: "[FATAL] ",
}

func write(level clog.LEVEL, skip int, format string, v ...interface{}) string {
	msg := ""
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		fn := runtime.FuncForPC(pc)
		var fnName string
		if fn == nil {
			fnName = "?()"
		} else {
			fnName = strings.TrimLeft(filepath.Ext(fn.Name()), ".") + "()"
		}
		msg = formats[level] + fmt.Sprintf("[%s:%d %s] ", file, line, fnName) + fmt.Sprintf(format, v...)
	}
	if len(msg) == 0 {
		msg = formats[level] + fmt.Sprintf(format, v...)
	}
	return msg
}
