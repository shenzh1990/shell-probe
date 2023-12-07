package util

import (
	"net"
	"time"
)

func CheckServer(addr string) (uint64, error) {
	timeout := time.Duration(5 * time.Second)
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", addr, timeout)
	delay := time.Now().Sub(t1)
	if err != nil {
		return 0, err
	}
	return uint64(delay), nil
}
