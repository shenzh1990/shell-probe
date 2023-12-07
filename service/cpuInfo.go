package service

import (
	"github.com/shenzh1990/shell-probe/model"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

// CPUInfo
//
//	@Description: 获取CPU信息
//	@return model.CPU
func CPUInfo() model.CPU {
	//获取CPU数据
	var cpuInfo model.CPU
	cpuInfo.CpuPhysicalCount, _ = cpu.Counts(false)
	cpuInfo.CpuLogicalCount, _ = cpu.Counts(true)
	cpuInfo.CpuTotalPercent, _ = cpu.Percent(1*time.Second, false)
	cpuInfo.CpuPerPercent, _ = cpu.Percent(1*time.Second, true)
	return cpuInfo
}
