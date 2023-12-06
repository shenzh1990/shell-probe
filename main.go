package main

import (
	"shell-probe/config"
	"shell-probe/router"
	"shell-probe/service"
	"strconv"
)

func main() {
	startParam := config.GetStartParam()
	switch startParam.M {
	case "wa":
		router.SetupRouter(strconv.Itoa(startParam.Ap))
	case "ws":
		service.StartWS(strconv.Itoa(startParam.Sp))
	case "all":
		go router.SetupRouter(strconv.Itoa(startParam.Ap))
		service.StartWS(strconv.Itoa(startParam.Sp))
	}
}
