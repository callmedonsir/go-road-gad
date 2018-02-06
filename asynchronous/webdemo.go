package main

import (
    "net/http"
    "strconv"
    "main/asynchronous/webdemo"
)
// 同步Handler
func syncHandler(w http.ResponseWriter, r *http.Request) {
    timer := webdemo.NewTimer("sync")
    defer timer.End()

    page := webdemo.NewSyncPage()
    for i := 0; i < 100; i ++ {
        key := strconv.Itoa(i)
        page.SetContents(key)
    }
    page.Render(w)
}
// 异步Handler
func asyncHandler(w http.ResponseWriter, r *http.Request) {
    timer := webdemo.NewTimer("async")
    defer timer.End()
    page := webdemo.NewAsyncPage()
    page.CountOut = 100
    for i:= 0; i < page.CountOut; i++ {
        key := strconv.Itoa(i)
        page.SetContents(key)
    }
    page.Render(w)
}

func main() {
    http.HandleFunc("/sync", syncHandler)
    http.HandleFunc("/async", asyncHandler)
    http.ListenAndServe("127.0.0.1:8888", nil)
}
