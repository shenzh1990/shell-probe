package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shenzh1990/shell-probe/util"
	"net/http"
)

// GetCPUInfo
//
//	@Description: 获取CPU信息
//	@param c:
func GetCPUInfo(c *gin.Context) {
	cpuInfo := CPUInfo()
	c.String(http.StatusOK, util.JsonResponse(0, util.SUCCESS_RTN, cpuInfo))
}

// GetMemInfo
//
//	@Description: 获取内存信息
//	@param c:
func GetMemInfo(c *gin.Context) {
	memInfo := MemInfo()
	c.String(http.StatusOK, util.JsonResponse(0, util.SUCCESS_RTN, memInfo))
}

// GetDiskInfo
//
//	@Description: 获取硬盘信息
//	@param c:
func GetDiskInfo(c *gin.Context) {
	diskInfo := DiskInfo()
	c.String(http.StatusOK, util.JsonResponse(0, util.SUCCESS_RTN, diskInfo))
}

// GetNetInfo
//
//	@Description: 获取网络信息
//	@param c:
func GetNetInfo(c *gin.Context) {
	netInfo := NetIO()
	c.String(http.StatusOK, util.JsonResponse(0, util.SUCCESS_RTN, netInfo))
}

// GetHostInfo
//
//	@Description: 响应HTTP以Get方式获取服务器硬件资源信息请求
//	@param c:
func GetHostInfo(c *gin.Context) {
	hostInfo := HostInfo()
	c.String(http.StatusOK, util.JsonResponse(0, util.SUCCESS_RTN, hostInfo))
}

// GetHostInfo
//
//	@Description: 响应HTTP以Get方式获取负载资源信息请求
//	@param c:
func GetLoadInfo(c *gin.Context) {
	loadInfo := LoadInfo()
	c.String(http.StatusOK, util.JsonResponse(0, util.SUCCESS_RTN, loadInfo))
}
