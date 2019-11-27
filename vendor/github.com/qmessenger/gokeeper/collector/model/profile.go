package model

import (
	"crypto/md5"
	"fmt"
)

// Profile数据的结构
type Profile struct {
	NodeId     string            `json:"NodeId"`
	Metric     string            `json:"Metric"`
	Tag        map[string]string `json:"Tag"`
	Section    string            `json:"Section"`
	RawSection string            `json:"RawSection"`
	Domain     string            `json:"Domain"`
	Component  string            `json:"Component"`
	Value      interface{}       `json:"Value"`
	Step       int               `json:"Step"`
	Timestamp  int64             `json:"Timestamp"`
}

// Profile结构主键
func (this *Profile) PK() string {
	return fmt.Sprintf("%s/%s/%s", this.NodeId, this.Metric, this.RawSection)
}

// Profile结构的MD5值
func (this *Profile) CheckSum() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(this.PK())))
}

// 查询参数结构
type ProfileQueryParam struct {
	NodeId    string `json:"NodeId"`
	Metric    string `json:"Metric"`
	Section   string `json:"Section"`
	Domain    string `json:"Domain"`
	Step      int    `json:"Step"`
	Start     int64  `json:"Start"`
	End       int64  `json:"End"`
	Path      string `json:"Path"`
	Timestamp int64  `json:"Timestamp"`
}

// 这个计算规则要和Profile.CheckSum()一致，因为要作为index的key，后面想个统一的规则？todo
func (this *ProfileQueryParam) PK() string {
	return fmt.Sprintf("%s/%s/%s", this.NodeId, this.Metric, this.Section)
}

func (this *ProfileQueryParam) CheckSum() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(this.PK())))
}

// 查询响应数据结构
type ProfileResponse struct {
	NodeId    string      `json:"NodeId"`
	Metric    string      `json:"Metric"`
	Timestamp int64       `json:"Timestamp"`
	Value     interface{} `json:"Values"`
}
