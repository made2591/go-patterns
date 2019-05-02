package logger

import (
	"fmt"
	"os"
	"sync"
)

var singletonInstance Logger
var mutex sync.Once

type Logger interface {
	Write(f *os.File, fr string, s interface{})
	SetLevel(l uint8)
	GetLevel() uint8
}

type logger struct {
	level uint8
}

func New() Logger {
	mutex.Do(func() {
		singletonInstance = &logger{}
	})
	return singletonInstance
}

func (l *logger) Write(f *os.File, fr string, s interface{}) {
	if f == nil {
		f = os.Stdout
	}
	fmt.Fprintf(f, fr, s)
}

func (l *logger) SetLevel(lv uint8) {
	l.level = lv
}

func (l *logger) GetLevel() uint8 {
	return l.level
}
