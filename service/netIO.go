package service

import (
	"github.com/shenzh1990/shell-probe/config"
	"github.com/shenzh1990/shell-probe/model"
	"github.com/shenzh1990/shell-probe/util"
	"github.com/shirou/gopsutil/net"
)

// NetIO
//
//	@Description: 获取网络IO信息
func NetIO() model.Net {
	//获取网络数据
	var netIO model.Net
	var delay model.Delay
	netInfo, _ := net.IOCounters(false)
	netIO.NetIO = netInfo

	delayTime, err := util.CheckServer(config.TestAddr)
	if err != nil {
		delay.IsConnected = 0
		delay.ErrorMessage = err.Error()
	} else {
		delay.IsConnected = 1
		delay.DelayTime = delayTime
	}
	netIO.Delay = delay
	return netIO
}
