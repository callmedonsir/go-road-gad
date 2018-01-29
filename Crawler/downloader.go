package Crawler

type Downloader interface {
	Module
	Download(req *Request)(*Response,error)
}