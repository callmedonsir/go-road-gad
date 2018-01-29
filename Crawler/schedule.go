package Crawler

import "net/http"

type Scheduler interface {
	Init(requestArgs RequestArgs,
		dataArgs DataArgs,
		moduleArgs ModuleArgs) (err error)
	Start(firstHttpReq *http.Request) (err error)
	Stop() (err error)
	Status() Status
	ErrorChan() <-chan error
	Idle() bool
	Summary() SchedSummary
}

type RequestArgs struct {
	AcceptedDomains []string `json:"accepted_primary_domains"`
	MaxDept         uint32   `json:"max_dept"`
}

type DataArgs struct {
	ReqBufferCap         uint32 `json:"req_buffer_cap"`
	ReqMaxBufferNumber   uint32 `json:"req_max_buffer_number"`
	RespBufferCap        uint32 `json:"resp_buffer_cap"`
	RespMapBufferNumber  uint32 `json:"resp_map_buffer_number"`
	ItemBufferCap        uint32 `json:"item_buffer_cap"`
	ItemMaxBufferNumber  uint32 `json:"item_max_buffer_number"`
	ErrorBufferCap       uint32 `json:"error_buffer_cap"`
	ErrorMaxBufferNumber uint32 `json:"error_max_buffer_number"`
}

type ModuleArgs struct {
	Downloader []Downloader
	Analyzers  []Analyzer
	Pipelines  []Pipeline
}

type Args interface {
	Check() error
}

type Status uint8

const (
	SCHED_STATUS_UNINITIALIZED Status = 0
	SCHED_STATUS_INITIALIZING  Status = 1
	SCHED_STATUS_INITIALIZED   Status = 2
	SCHED_STATUS_STARTING      Status = 3
	SCHED_STATUS_STARTED       Status = 4
	SCHED_STATUS_STOPPING      Status = 5
	SCHED_STATUS_STOPPED       Status = 6
)
