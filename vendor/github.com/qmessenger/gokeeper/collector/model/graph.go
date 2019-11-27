package model

import (
	"crypto/md5"
	"fmt"
)

// 绘图数据结构
type GraphItem struct {
	NodeId     string  `json:"NodeId"`
	Metric     string  `json:"Metric"`
	Section    string  `json:"Section"`
	RawSection string  `json:"RawSection"`
	Value      float64 `json:"Value"`
	Step       int     `json:"Step"`
	Type       string  `json:"RRDType"`
	Min        string  `json:"Min"`
	Max        string  `json:"Max"`
	Heartbeat  int     `json:"Heartbeat"`
	Timestamp  int64   `json:"Timestamp"`
}

// GraphItem主键
func (this *GraphItem) PK() string {
	return fmt.Sprintf("%s/%s/%s", this.NodeId, this.Metric, this.RawSection)
}

// GraphItem唯一MD5值
func (this *GraphItem) CheckSum() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(this.PK())))
}

// 官方文档里dsname格式定义：长度0-19位，且必须是[a-zA-Z0-9_]
func (this *GraphItem) NodeIdToDsName() string {
	s := fmt.Sprintf("%x", md5.Sum([]byte(this.NodeId)))
	return s[0:15]
}

// 查询rrd返回的数据结构
type RRDData struct {
	Timestamp int64   `json:"Timestamp"`
	Value     float64 `json:"Value"`
}

// 查询参数结构
type GraphQueryParam struct {
	NodeId  string `json:"NodeId"`
	Metric  string `json:"Metric"`
	CF      string `json:"CF"`
	Section string `json:"Section"`
	Step    int    `json:"Step"`
	Start   int64  `json:"Start"`
	End     int64  `json:"End"`
}

// 这个计算规则要和GraphItem.CheckSum()一致，因为要作为index的key，后面想个统一的规则？todo
func (this *GraphQueryParam) PK() string {
	return fmt.Sprintf("%s/%s/%s", this.NodeId, this.Metric, this.Section)
}

func (this *GraphQueryParam) CheckSum() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(this.PK())))
}

// 查询响应数据结构
type GraphResponse struct {
	NodeId string      `json:"NodeId"`
	Metric string      `json:"Metric"`
	Step   int         `json:"Step"`
	Values interface{} `json:"Values"`
}

// 生成graph时Dsname对应所需的信息
type GraphQueryDsName struct {
	DsName  string
	RRDFile string
}

func NewGraphQueryDsName(dsName, rrdFile string) *GraphQueryDsName {
	return &GraphQueryDsName{DsName: dsName, RRDFile: rrdFile}
}
