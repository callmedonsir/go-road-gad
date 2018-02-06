package webdemo

import (
    "time"
    "fmt"
)

const (
    // 休眠 20,000,000 ns，用以模拟数据I/O的延迟
    SLEEP = 200000000
)

// 取数据
func GetContents(key string) string {
    time.Sleep(SLEEP)
    return fmt.Sprintf("%s. The quick brown fox jumps over the lazy dog.", key)
}
