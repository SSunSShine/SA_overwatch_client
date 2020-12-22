package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

type mInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	ModelName   string  `json:"model_name"`
}

func MemInfo(c *gin.Context) {
	memory, _ := mem.VirtualMemory()
	var info mInfo
	info.Total = memory.Total
	info.Available = memory.Free + memory.Available
	info.Used = memory.Used
	info.UsedPercent = memory.UsedPercent
	cpuInfo, _ := cpu.Info()
	info.ModelName = cpuInfo[0].ModelName

	c.JSON(http.StatusOK, gin.H{
		"msg": info,
	})
}
