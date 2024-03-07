package sharp

import (
	"github.com/natefinch/lumberjack"
	"io"
)

func NewFileRotateWriter(file string, maxSize, maxBackups, maxAge int) io.Writer {
	return &lumberjack.Logger{
		Filename:   file,
		MaxSize:    maxSize,    // megabytes
		MaxBackups: maxBackups, // how many files will be kept
		MaxAge:     maxAge,     // how long the files will be kept (days)
		Compress:   false,
	}
}
