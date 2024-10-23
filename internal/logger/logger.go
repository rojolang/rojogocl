package logger

import (
    "sync"
)

type Logger struct {
    verbosity int
    mu        sync.Mutex
}

func NewLogger(verbosity int) *Logger {
    return &Logger{
        verbosity: verbosity,
    }
}

func (l *Logger) VerbosePrint(msg string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    // TODO: Implement logging
}
