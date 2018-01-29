package Crawler

import "net/http"

type Request struct {
	//http请求
	httpReq *http.Request
	//请求的深度
	depth uint32
}

type Response struct {
	httpResp *http.Response
	depth    uint32
}

type Item map[string]interface{}

type Data interface {
	Valid() bool
}

func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq,
		depth: depth}
}

func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

func (req *Request) Depth() uint32 {
	return req.depth
}

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp,
		depth: depth}
}

func (resp *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

func (resp *Response) Depth() uint32 {
	return resp.depth
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

func (item Item) Valid() bool {
	return item != nil
}

type MID string

type Module interface {
	//用于获取当前组件的ID
	ID() MID
	//用于获取当前组件的网络地址
	Addr() string
	//用于获取当前组件的评分
	Score() uint64
	//设置评分
	SetScore(score uint64)
	//获取评分计算器
	ScoreCalculator() CalculateScore
	//获取当前组件被调用次数
	CalledCount() uint64
	//获取当前组件接受的调用的次数
	//组件一般会由于超负荷或参数有误而拒绝调用
	AcceptedCount() uint64
	//获取当前组件已经成功完成调用的次数
	CompletedCount() uint64
	//获取当前组件正在处理的调用的次数
	HandlingNumber() uint64
	//用于一次性获取所有次数
	Counts() Counts
	//获取组件摘要
	Summary() SummaryStruct
}

type Type string

const (
	//下载器
	TYPE_DOWNLOADER Type = "downloader"
	//分析器
	TYPE_ANALYZER Type = "analyzer"
	//条目处理管道
	TYPE_PIPELINE Type = "pipeline"
)

//合法的组件类型-字母的映射
var legalTypeLetterMap = map[Type]string{
	TYPE_DOWNLOADER: "D",
	TYPE_ANALYZER:   "A",
	TYPE_PIPELINE:   "P",
}

//序列号生成器接口类型
type SNGenertor interface {
	Start() uint64
	Max() uint64
	Next() uint64
	CycleCount() uint64
	Get() uint64
}

//组件注册器
type Registrar interface {
	//用于注册组件实例
	Registrar(moudle Moudle) (bool, error)
	//用于注册组件势力
	Unregister(mid MID) (bool, error)
	//获取一个指定类型组件：基于负载均衡策略
	Get(moudleType Type) (Moudle, error)
	//获取指定类型的所有组件
	GetAllByType(moudleType Type) (map[MID]Moudle, error)
	//获取所有组件的实例
	GetAll() map[MID]Moudle
	//清楚所有组件注册记录
	Clear()
}

type CalculateScore func(counts Counts) uint64

type Counts struct{}

type SummaryStruct struct {
	ID        MID         `json:"id"`
	Called    uint64      `json:"called"`
	Accepet   uint64      `json:"accepet"`
	Completed uint64      `json:"completed"`
	Handling  uint64      `json:"handling"`
	Extra     interface{} `json:"extra"`
}

