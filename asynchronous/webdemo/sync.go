package webdemo

import (
    "net/http"
    "fmt"
    "strconv"
)

type SyncContents map[string]string

type SyncPage struct {
    contents SyncContents
    page Page
}

func NewSyncPage() *SyncPage {
    return &SyncPage{make(SyncContents), nil}
}

func (page *SyncPage) SetContents(key string) {
    page.contents[key] = GetContents(key)
}

func (page *SyncPage) Render(w http.ResponseWriter) {
    lines := ""
    for i := 0; i < len(page.contents); i++ {
        key := strconv.Itoa(i)
        lines += fmt.Sprintf(TEMPLATE_LINE, page.contents[key])
    }
    block := fmt.Sprintf(TEMPLATE_BLOCK, lines)
    fmt.Fprintf(w, TEMPLATE_PAGE, block)
}
