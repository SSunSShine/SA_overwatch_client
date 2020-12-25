package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/process"
	"net/http"
)

type ProcInfo struct {
	Pid        int32   `json:"pid"`
	Name       string  `json:"name"`
	CreateTime int64   `json:"create_time"`
	CpuPercent float64 `json:"cpu_percent"`
	MemInfo    float64 `json:"mem_info"`
}

func ProcsInfo(c *gin.Context)  {
	var procs []ProcInfo
	processes, _ := process.Processes()
	for _, p := range processes {
		var proc ProcInfo
		proc.Pid = p.Pid
		proc.Name, _ = p.Name()
		proc.CreateTime, _ = p.CreateTime()
		cpuPercent, _ := p.CPUPercent()
		proc.CpuPercent = cpuPercent * 100
		memInfo, _ := p.MemoryInfo()
		if memInfo != nil && memInfo.RSS != 0 {
			proc.MemInfo = float64(memInfo.RSS/(1024*1024))
		}
		procs = append(procs, proc)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": procs,
	})
}
