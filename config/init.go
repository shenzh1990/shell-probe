package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	version  string
	h        bool
	v, V     bool
	TestAddr string
	Ap       int
	Sign     string
	FCpu     bool
)

func usage() {
	fmt.Fprintf(os.Stderr, `server resource probe version: `+version+`
Usage: server resource probe  [-ap WebApiServer start port default[8082]] [-token set visit password defaut[123456] [-ta set test net delay addr defaut[127.0.0.1:8082]] [-fc hostinfo contains cpu info defaut[false]]
Options:
`)
	flag.PrintDefaults()
}
func InitConfig() {
	version = "0.0.1"
	flag.BoolVar(&h, "h", false, "show help")
	flag.StringVar(&Sign, "token", "123456", "set api password")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and exit")
	flag.BoolVar(&FCpu, "fc", false, "show cpuinfo in host info ")
	flag.IntVar(&Ap, "ap", 8082, "WebApiServer start port")
	flag.StringVar(&TestAddr, "ta", "127.0.0.1:8082", "set test net delay addr")
	flag.Parse()
	if h {
		flag.Usage = usage
		flag.Usage()
		os.Exit(0)
	} else if v || V {
		fmt.Fprintf(os.Stderr, `server resource probe version: `+version)
		os.Exit(0)
	}
}
