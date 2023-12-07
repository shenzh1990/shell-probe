package service

import (
	"github.com/shenzh1990/shell-probe/config"
	"github.com/shenzh1990/shell-probe/model"
)

// HostInfo
//
//	@Description: 获取服务器硬件资源信息函数
//	@param c:
func HostInfo() model.HostInfo {
	var hostInfo model.HostInfo
	//获取CPU数据
	if config.FCpu {
		hostInfo.CPU = CPUInfo()
	}
	//获取内存数据
	hostInfo.Mem = MemInfo()
	//获取硬盘数据
	hostInfo.DiskInfo = DiskInfo()
	//获取网络数据
	hostInfo.Net = NetIO()
	//获取负载信息
	hostInfo.Load = LoadInfo()
	return hostInfo

}
