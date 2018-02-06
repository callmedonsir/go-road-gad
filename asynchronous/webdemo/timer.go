package webdemo

import (
    "log"
    "time"
)

type Timer struct {
    t int64
    name string
}

func NewTimer(name string) *Timer{
    return &Timer{int64(time.Nanosecond), name}
}

func (timer *Timer) End() {
    log.Printf("%s:\t%d", timer.name, int64(time.Nanosecond) - timer.t)
}
