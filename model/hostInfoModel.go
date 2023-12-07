package model

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
)

// CPU
// @Description: CPU信息结构体
type CPU struct {
	CpuPhysicalCount int       `json:"cpu_physical_count"` //物理核心数
	CpuLogicalCount  int       `json:"cpu_logical_count"`  //逻辑核心数
	CpuTotalPercent  []float64 `json:"cpu_total_percent"`  //CPU总占用
	CpuPerPercent    []float64 `json:"cpu_per_percent"`    //每一个逻辑核心的占用
}

// Mem
// @Description: 内存信息结构体
type Mem struct {
	MemTotal       uint64  `json:"mem_total"`        //内存总量Gb
	MemAvailable   uint64  `json:"men_available"`    //内存可用容量Gb
	MemUsedPercent float64 `json:"mem_used_percent"` //内存已使用百分比
}

// Disk
// @Description: 硬盘信息结构体
type Disk struct {
	Driver string              `json:"driver"`
	Usage  *disk.UsageStat     `json:"disk_usage"`
	DiskIO disk.IOCountersStat `json:"disk_io"`
}
type DiskInfo struct {
	DiskTotal       uint64  `json:"disk_total"`
	DiskFree        uint64  `json:"disk_free"`
	DiskUsed        uint64  `json:"disk_used"`
	DiskUsedPercent float64 `json:"usedPercent"`
	DiskList        []Disk  `json:"disk_list"`
}

// Net
// @Description: 网络信息结构体
type Net struct {
	NetIO []net.IOCountersStat `json:"net_io"` //网络IO
	Delay Delay                `json:"delay"`  //网络IO
}

type Delay struct {
	IsConnected  int    `json:"is_connected"` //1通 0不通
	DelayTime    uint64 `json:"delay_time"`   //单位ns纳秒 需自己转成ms
	ErrorMessage string `json:"error_message"`
}

// load
// @Description: load信息结构体
type Load struct {
	LoadAvg load.AvgStat `json:"load_avg"` //负载平均
}

// HostInfo
// @Description: 主机信息结构体
type HostInfo struct {
	CPU      `json:"cpu"`
	Mem      `json:"mem"`
	DiskInfo `json:"disk"`
	Net      `json:"net"`
	Load     `json:"load"`
}
