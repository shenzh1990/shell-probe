package main

import (
	"github.com/shenzh1990/shell-probe/config"
	"github.com/shenzh1990/shell-probe/router"
)

func main() {
	config.InitConfig()
	router.SetupRouter(config.Ap)
	//delayTime, err := util.CheckServer(config.TestAddr)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Println(delayTime)
}
