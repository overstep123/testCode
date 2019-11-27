package model

import (
	"sync"
)

// 配置的结构
type Config struct {
	sync.RWMutex
	Base    *BaseConfig
	GoRoot  map[string]*GorootConfig
	Metrics map[string]*MetricConfig
	Nodeids map[string]map[string]*NodeidsConfig
}

func NewConfig() *Config {
	return &Config{
		Base:    new(BaseConfig),
		GoRoot:  make(map[string]*GorootConfig),
		Metrics: make(map[string]*MetricConfig),
		Nodeids: make(map[string]map[string]*NodeidsConfig),
	}
}

type BaseConfig struct {
	LogPath   string   `ini:"log_path"`
	RRDPath   string   `ini:"rrd_path"`
	PngPath   string   `ini:"png_path"`
	BinPath   string   `ini:"bin_path"`
	ProfPath  string   `ini:"prof_path"`
	GraphPath string   `ini:"graph_path"`
	Collect   []string `ini:"collect"`
}

type GorootConfig struct {
	GoRoot string `ini:"goroot"`
}

type MetricConfig struct {
	Step    int      `ini:"step"`
	Domains []string `ini:"domains"`
}

type NodeidsConfig struct {
	Nodeids []string `ini:"nodeids"`
}
