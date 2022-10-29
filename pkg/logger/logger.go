package logger

import (
	"io"
	"log"
	"os"
	"sync/atomic"
)

// log level
const (
	LDEBUG = iota + 1 // 1
	LWARN             // 2
	LINFO             // 3
	LERROR            // 4
	LFATAL            // 5
)

type MyLogger struct {
	level       int64
	w           io.Writer
	debugLogger *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
}

func New(w io.Writer, level int64, flag int) *MyLogger {
	if w == nil {
		w = os.Stderr
	}

	if flag <= 0 {
		flag = log.LstdFlags
	}

	return &MyLogger{
		w:           w,
		level:       level,
		debugLogger: log.New(w, "[DEBUG] ", flag|log.Lmsgprefix),
		warnLogger:  log.New(w, "[WARN] ", flag|log.Lmsgprefix),
		infoLogger:  log.New(w, "[INFO] ", flag|log.Lmsgprefix),
		errLogger:   log.New(w, "[ERROR] ", flag|log.Lmsgprefix),
		fatalLogger: log.New(w, "[FATAL] ", flag|log.Lmsgprefix),
	}
}

func (l *MyLogger) SetLevel(level int64) {
	if level < LDEBUG || level > LFATAL {
		return
	}

	atomic.StoreInt64(&l.level, level)
}

func (l *MyLogger) Debugln(v ...interface{}) {
	if atomic.LoadInt64(&l.level) > LDEBUG {
		return
	}
	l.debugLogger.Println(v...)
}

func (l *MyLogger) Debugf(format string, v ...interface{}) {
	if atomic.LoadInt64(&l.level) > LDEBUG {
		return
	}
	l.debugLogger.Printf(format, v...)
}

func (l *MyLogger) Infoln(v ...interface{}) {
	if atomic.LoadInt64(&l.level) > LINFO {
		return
	}
	l.infoLogger.Println(v...)
}

func (l *MyLogger) Infof(format string, v ...interface{}) {
	if atomic.LoadInt64(&l.level) > LINFO {
		return
	}
	l.infoLogger.Printf(format, v...)
}
func (l *MyLogger) PF(lv int, format string, v ...interface{}) {
	switch lv {
	case LERROR:
		if atomic.LoadInt64(&l.level) > LERROR {
			return
		}
		l.errLogger.Printf(format, v...)
	case LDEBUG:
		if atomic.LoadInt64(&l.level) > LDEBUG {
			return
		}
		l.debugLogger.Printf(format, v...)
	case LFATAL:
		if atomic.LoadInt64(&l.level) > LDEBUG {
			return
		}
		l.fatalLogger.Printf(format, v...)
	case LWARN:
		if atomic.LoadInt64(&l.level) > LDEBUG {
			return
		}
		l.warnLogger.Printf(format, v...)
	case LINFO:
		if atomic.LoadInt64(&l.level) > LINFO {
			return
		}
		l.infoLogger.Printf(format, v...)
	}

}
