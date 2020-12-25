package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/disk"
	"net/http"
)

type DiskInfo struct {
	Device      string  `json:"device"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func Disk(c *gin.Context)  {
	var infos []DiskInfo
	partitions, _ := disk.Partitions(true)
	for _, p := range partitions {
		var info DiskInfo
		info.Device = p.Device
		usage, _ := disk.Usage(p.Device)
		info.Total = usage.Total
		info.Used = usage.Used
		info.Free = usage.Free
		info.UsedPercent = usage.UsedPercent
		infos = append(infos, info)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": infos,
	})
}