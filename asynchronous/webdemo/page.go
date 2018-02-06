package webdemo

import (
    "net/http"
)

const (
    TEMPLATE_PAGE = "<html><head></head><body>%s</body></html>"
    TEMPLATE_BLOCK = "<ul>%s</ul>"
    TEMPLATE_LINE = "<li>%s</li>"
)

type Page interface {
    SetContents(key string)
    Render(w http.ResponseWriter)
}
