package webdemo

import (
    "net/http"
    "fmt"
    "time"
)

const (
    // 超时 500,000,000 ns
    TIMEOUT = 500000000
)
// 页面内容
type AsyncContents struct {
    key, value string
}
// 页面
type AsyncPage struct {
    contents chan AsyncContents
    timeout chan bool
    CountOut int
    page Page
}

func NewAsyncPage() *AsyncPage {
    return &AsyncPage{make(chan AsyncContents), make(chan bool), 0, nil}
}

func (page *AsyncPage) SetContents(key string) {
    // 异步的数据获取
    go func() {
        page.contents <- AsyncContents{key, GetContents(key)}
    }()
    // 设置针对页面的超时
    go func() {
        time.Sleep(TIMEOUT)
        page.timeout <- true
    }()
}

func (page *AsyncPage) Render(w http.ResponseWriter) {

    lines := ""
    LOOP: for i := 0; i < page.CountOut; i++{
        select {
            case line := <-page.contents:
                lines = fmt.Sprintf("%s" + TEMPLATE_LINE, lines, line.value)
                // 每获取一个数据，就去掉一个超时
                go func() {<-page.timeout}()
            case <-page.timeout:
                lines = fmt.Sprintf("%s" + TEMPLATE_LINE, lines, "Time Out")
                break LOOP
        }
    }
    block := fmt.Sprintf(TEMPLATE_BLOCK, lines)
    fmt.Fprintf(w, TEMPLATE_PAGE, block)
}
