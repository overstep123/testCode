package model

import (
	"crypto/md5"
	"encoding/gob"
	"fmt"
)

func init() {
	gob.RegisterName(PkgPrefix + "EventMetric", EventMetric{})
	gob.RegisterName(PkgPrefix + "MetricValue", MetricValue{})
	gob.Register([]*MetricValue{})
}

const PkgPrefix string = "github.com/qmessenger/gokeeper/collector/model."

//采集项数据
type MetricValue struct {
	NodeId    string      `json:"NodeId"`
	Metric    string      `json:"Metric"`
	Tag       string      `json:"Tag"`
	Section   string      `json:"Section"`
	Value     interface{} `json:"Value"`
	Step      int         `json:"Step"`
	Type      string      `json:"RrdType"`
	Timestamp int64       `json:"Timestamp"`
	Extra     interface{} `json:"Extra"`
}

func NewMetricValue(metric string, value interface{}, mtype, section string) *MetricValue {
	return &MetricValue{
		Metric:  metric,
		Value:   value,
		Type:    mtype,
		Section: section,
	}
}

// 初始化GAUGE类型metric结构
func GaugeValue(metric string, value ...interface{}) *MetricValue {
	if len(value) == 2 {
		if section, ok := value[1].(string); ok {
			return NewMetricValue(metric, value[0], "GAUGE", section)
		}
	}
	return NewMetricValue(metric, value[0], "GAUGE", "")
}

// 采集项主键
func (this *MetricValue) PK() string {
	return fmt.Sprintf("%s/%s/%s", this.NodeId, this.Metric, this.Section)
}

// 采集项结构MD5值
func (this *MetricValue) CheckSum() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(this.PK())))
}

// 给node发送采集事件时，传输的data数据结构
type EventMetric struct {
	Step int `json:"Step"`
}
