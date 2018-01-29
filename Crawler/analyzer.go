package Crawler

import "net/http"

type Analyzer interface {
	Module
	//用于返回钱分析器使用的相应解析函数的列表
	RespParsers()[]ParseResponse
	Analyz(resp *Response)([]Data,[]error)
}

type ParseResponse func(httpResp *http.Response,respDepth uint32)([]Data,[]error)
