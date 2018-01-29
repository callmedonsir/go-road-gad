package Crawler

import (
	"strings"
	"bytes"
	"fmt"
)

type ErrorType string

const (
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	ERROR_TYPE_ANALYZER   ErrorType = "analyzer error"
	ERROR_TYPE_PIPELINE   ErrorType = "pipeline error"
	ERROR_TYPE_SCHEDULER  ErrorType = "scheduler error"
)


type CrawlerError interface {
	Type() ErrorType
	Error() string
}

type myCrawlerError struct {
	errType    ErrorType
	errMsg     string
	fullErrMsg string
}

func NewCrawlerError(errType ErrorType,errMsg string) CrawlerError{
	return &myCrawlerError{
		errType:errType,
		errMsg:strings.TrimSpace(errMsg),
	}
}

func (ce *myCrawlerError)Type()ErrorType{
	return ce.errType
}

func (ce *myCrawlerError)Error()string{
	if ce.fullErrMsg==""{
		return ce.errMsg
	}
	return ce.fullErrMsg
}

func (ce *myCrawlerError)getFullErrMsg(){
	var buffer bytes.Buffer
	buffer.WriteString("crawler error:")
	if ce.errType!=""{
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf(buffer.String())
	return
}
