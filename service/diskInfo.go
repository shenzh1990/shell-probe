package service

import (
	"github.com/shenzh1990/shell-probe/model"
	"github.com/shirou/gopsutil/disk"
)

// DiskInfo
//
//	@Description: 获取硬盘信息列表
func DiskListInfo() []model.Disk {
	//获取硬盘数据
	var diskList []model.Disk
	info, _ := disk.Partitions(true) //所有分区

	info3, _ := disk.IOCounters() //所有硬盘的io信息
	for _, device := range info {
		var diskTemp model.Disk
		diskTemp.Driver = device.Device
		info2, _ := disk.Usage(device.Device) //指定某路径的硬盘使用情况
		diskTemp.Usage = info2
		diskTemp.DiskIO = info3[device.Device]
		diskList = append(diskList, diskTemp)
	}
	return diskList
}
func DiskInfo() model.DiskInfo {
	//获取硬盘数据
	var diskInfo model.DiskInfo
	var diskList []model.Disk
	var diskTotal uint64
	var diskFree uint64
	var diskUsed uint64
	var diskUsedPercent float64      // `json:"usedPercent"`
	info, _ := disk.Partitions(true) //所有分区
	info3, _ := disk.IOCounters()    //所有硬盘的io信息
	for _, device := range info {
		var diskTemp model.Disk
		diskTemp.Driver = device.Device
		info2, _ := disk.Usage(device.Mountpoint) //指定某路径的硬盘使用情况
		diskTemp.Usage = info2
		diskTemp.DiskIO = info3[device.Mountpoint]
		diskList = append(diskList, diskTemp)
		if info2 != nil && device.Mountpoint == "/" {
			diskTotal = info2.Total
			diskFree = info2.Free
			diskUsed = info2.Used
			diskUsedPercent = info2.UsedPercent
		}

	}
	diskInfo.DiskList = diskList
	diskInfo.DiskTotal = diskTotal
	diskInfo.DiskUsed = diskUsed
	diskInfo.DiskFree = diskFree
	diskInfo.DiskUsedPercent = diskUsedPercent
	return diskInfo
}
