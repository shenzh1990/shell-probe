package service

import (
	"github.com/shenzh1990/shell-probe/model"
	"github.com/shirou/gopsutil/mem"
)

func MemInfo() model.Mem {
	//获取内存数据
	var m model.Mem
	v, _ := mem.VirtualMemory()
	m.MemTotal = v.Total
	m.MemAvailable = v.Available
	m.MemUsedPercent = v.UsedPercent
	return m
}
