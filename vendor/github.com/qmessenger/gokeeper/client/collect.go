package client

import (
	"bytes"
	"runtime"
	"runtime/pprof"
	"time"

	cmodel "github.com/qmessenger/gokeeper/collector/model"
	"github.com/qmessenger/gokeeper/model"
)

// 回调方法列表
type collectFuncs struct {
	Fs []func() []*cmodel.MetricValue
}

var collectMappers map[int]collectFuncs

// 初始化事件回调mapper
func init() {
	collectMappers = map[int]collectFuncs{
		model.EventCollectProc: collectFuncs{
			Fs: []func() []*cmodel.MetricValue{
				CpuMetrics,
				MemMetrics,
			},
		},
		model.EventCollectPProf: collectFuncs{
			Fs: []func() []*cmodel.MetricValue{
				GoroutineMetric,
				AllocMetric,
				HeapMetric,
				StackMetric,
				GcMetric,
			},
		},
		model.EventCollectProfile: collectFuncs{
			Fs: []func() []*cmodel.MetricValue{
				ProfilingHeapMetric,
				ProfilingCpuMetric,
				ProfilingGoroutineMetric,
				ProfilingThreadcreateMetric,
				ProfilingBlockMetric,
			},
		},
	}
}

// 执行数据采集
func collect(c *Client, fns []func() []*cmodel.MetricValue, metric cmodel.EventMetric) []*cmodel.MetricValue {
	mvs := []*cmodel.MetricValue{}
	for _, fn := range fns {
		item := fn()
		for _, mv := range item {
			mvs = append(mvs, mv)
		}
	}
	ts := time.Now().Unix()
	mvsLen := len(mvs)
	for i := 0; i < mvsLen; i++ {
		mvs[i].NodeId = c.node.GetID()
		mvs[i].Step = metric.Step
		mvs[i].Timestamp = ts
		mvs[i].Extra = c.node
	}
	return mvs
}

// 进程cpu数据
func CpuMetrics() []*cmodel.MetricValue {
	ProcInfo.Cpu.Refresh()
	ProcInfo.Cpu.CurrentUsage()
	return []*cmodel.MetricValue{cmodel.GaugeValue("cpu.usage", ProcInfo.Cpu.CpuUsage)}
}

// 进程内存数据
func MemMetrics() []*cmodel.MetricValue {
	ProcInfo.Mem.Refresh()
	virt := cmodel.GaugeValue("mem.virt", ProcInfo.Mem.VmSize)
	res := cmodel.GaugeValue("mem.res", ProcInfo.Mem.VmRss)
	return []*cmodel.MetricValue{virt, res}
}

// 进程协程数量
func GoroutineMetric() []*cmodel.MetricValue {
	return []*cmodel.MetricValue{cmodel.GaugeValue("go.goroutines", runtime.NumGoroutine())}
}

// 进程内存申请数据
func AllocMetric() []*cmodel.MetricValue {
	var memoryStats runtime.MemStats
	runtime.ReadMemStats(&memoryStats)
	alloc := cmodel.GaugeValue("go.Alloc", memoryStats.Alloc)
	totalAlloc := cmodel.GaugeValue("go.TotalAlloc", memoryStats.TotalAlloc)
	sys := cmodel.GaugeValue("go.Sys", memoryStats.Sys)
	lookups := cmodel.GaugeValue("go.Lookups", memoryStats.Lookups)
	mallocs := cmodel.GaugeValue("go.Mallocs", memoryStats.Mallocs)
	frees := cmodel.GaugeValue("go.Frees", memoryStats.Frees)

	return []*cmodel.MetricValue{alloc, totalAlloc, sys, lookups, mallocs, frees}
}

// 进程heap数据
func HeapMetric() []*cmodel.MetricValue {
	var memoryStats runtime.MemStats
	runtime.ReadMemStats(&memoryStats)
	heapAlloc := cmodel.GaugeValue("go.HeapAlloc", memoryStats.HeapAlloc)
	heapSys := cmodel.GaugeValue("go.HeapSys", memoryStats.HeapSys)
	heapIdle := cmodel.GaugeValue("go.HeapIdle", memoryStats.HeapIdle)
	heapInuse := cmodel.GaugeValue("go.HeapInuse", memoryStats.HeapInuse)
	heapReleased := cmodel.GaugeValue("go.HeapReleased", memoryStats.HeapReleased)
	heapObjects := cmodel.GaugeValue("go.HeapObjects", memoryStats.HeapObjects)

	return []*cmodel.MetricValue{heapAlloc, heapSys, heapIdle, heapInuse, heapReleased, heapObjects}
}

// 进程stack数据
func StackMetric() []*cmodel.MetricValue {
	var memoryStats runtime.MemStats
	runtime.ReadMemStats(&memoryStats)
	stackInuse := cmodel.GaugeValue("go.StackInuse", memoryStats.StackInuse)
	mspanInuse := cmodel.GaugeValue("go.MSpanInuse", memoryStats.MSpanInuse)
	mspanSys := cmodel.GaugeValue("go.MSpanSys", memoryStats.MSpanSys)
	mcacheInuse := cmodel.GaugeValue("go.MCacheInuse", memoryStats.MCacheInuse)
	mcacheSys := cmodel.GaugeValue("go.MCacheSys", memoryStats.MCacheSys)
	buckHashSys := cmodel.GaugeValue("go.BuckHashSys", memoryStats.BuckHashSys)

	return []*cmodel.MetricValue{stackInuse, mspanInuse, mspanSys, mcacheInuse, mcacheSys, buckHashSys}
}

// 进程GC数据
func GcMetric() []*cmodel.MetricValue {
	var memoryStats runtime.MemStats
	runtime.ReadMemStats(&memoryStats)
	nextGC := cmodel.GaugeValue("go.NextGC", memoryStats.NextGC)
	lastGC := cmodel.GaugeValue("go.LastGC", memoryStats.LastGC)
	pauseTotalNs := cmodel.GaugeValue("go.PauseTotalNs", memoryStats.PauseTotalNs)
	pauseNs := cmodel.GaugeValue("go.PauseNs", memoryStats.PauseNs[(memoryStats.NumGC+255)%256])
	numGC := cmodel.GaugeValue("go.NumGC", memoryStats.NumGC)

	return []*cmodel.MetricValue{nextGC, lastGC, pauseTotalNs, pauseNs, numGC}
}

// 进程profile heap数据
func ProfilingHeapMetric() []*cmodel.MetricValue {
	var buff bytes.Buffer
	profile := pprof.Lookup("heap")
	profile.WriteTo(&buff, 1)
	heap := cmodel.GaugeValue("profiling.heap", buff.String())
	return []*cmodel.MetricValue{heap}
}

// 进程profile goroutine数据
func ProfilingGoroutineMetric() []*cmodel.MetricValue {
	var buff bytes.Buffer
	profile := pprof.Lookup("goroutine")
	profile.WriteTo(&buff, 1)
	goroutine := cmodel.GaugeValue("profiling.goroutine", buff.String())
	return []*cmodel.MetricValue{goroutine}
}

// 进程profile threadcreate数据
func ProfilingThreadcreateMetric() []*cmodel.MetricValue {
	var buff bytes.Buffer
	profile := pprof.Lookup("threadcreate")
	profile.WriteTo(&buff, 1)
	threadcreate := cmodel.GaugeValue("profiling.threadcreate", buff.String())
	return []*cmodel.MetricValue{threadcreate}
}

// 进程profile block数据
func ProfilingBlockMetric() []*cmodel.MetricValue {
	var buff bytes.Buffer
	profile := pprof.Lookup("block")
	profile.WriteTo(&buff, 1)
	block := cmodel.GaugeValue("profiling.block", buff.String())
	return []*cmodel.MetricValue{block}
}

// 进程 profile cpu数据
func ProfilingCpuMetric() []*cmodel.MetricValue {
	var buff bytes.Buffer
	pprof.StartCPUProfile(&buff)
	time.Sleep(time.Duration(30) * time.Second)
	pprof.StopCPUProfile()
	cpu := cmodel.GaugeValue("profiling.cpu", buff.String())
	return []*cmodel.MetricValue{cpu}
}
