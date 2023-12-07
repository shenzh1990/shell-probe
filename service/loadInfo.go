package service

import (
	"github.com/shenzh1990/shell-probe/model"
	log "github.com/sirupsen/logrus"

	"github.com/shirou/gopsutil/load"
)

// Load
//
//	@Description: 获取负载信息
//	@return model.Load
func LoadInfo() model.Load {
	var loadInfo model.Load
	var loadAvg *load.AvgStat

	loadAvg, err := load.Avg()
	if err != nil {
		log.Println("load.Avg():", err)
	}
	loadInfo.LoadAvg = *loadAvg
	return loadInfo
}
