package telemetry

import (
	"runtime"
	"sync"
	"time"
)


type SystemUsage struct {
	Memory     memoryUsage
	CpuBusyAvg float64
}

type memoryUsage struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects uint64
}

type CpuUsage struct {
	Busy, // time used by all processes. this ideally does not include system processes.
	Idle, // time used by the idle process
	Total uint64 // reported sum total of all usage
}

var once sync.Once
var lastSample CpuUsage
var usageAvg float64

func NewSystemUsage() (s SystemUsage) {
	// The micro-service is to be considered the System Of Record (SOR) in terms of accurate information.
	// Fetch metrics for the metadata service.
	var rtm runtime.MemStats

	// Read full memory stats
	runtime.ReadMemStats(&rtm)

	// Miscellaneous memory stats
	s.Memory.Alloc = rtm.Alloc
	s.Memory.TotalAlloc = rtm.TotalAlloc
	s.Memory.Sys = rtm.Sys
	s.Memory.Mallocs = rtm.Mallocs
	s.Memory.Frees = rtm.Frees

	// Live objects = Mallocs - Frees
	s.Memory.LiveObjects = s.Memory.Mallocs - s.Memory.Frees

	s.CpuBusyAvg = usageAvg

	return s
}

func StartCpuUsageAverage() {
	once.Do(func() {
		for {
			nextUsage := PollCpu()
			usageAvg = AvgCpuUsage(lastSample, nextUsage)
			lastSample = nextUsage

			time.Sleep(time.Second * 10)
		}
	})
}


func PollCpu() (cpuSnapshot CpuUsage) {
	return cpuSnapshot
}

func AvgCpuUsage(init, final CpuUsage) (avg float64) {
	return -1
}
