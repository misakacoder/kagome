package net

import (
	"github.com/misakacoder/kagome/errs"
	"net"
)

const Localhost = "127.0.0.1"

func GetLocalAddr() (addr []string) {
	addr = append(addr, Localhost)
	interfaces, err := net.Interfaces()
	errs.Panic(err)
	for _, value := range interfaces {
		if (value.Flags & net.FlagUp) != 0 {
			addresses, _ := value.Addrs()
			for _, address := range addresses {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						addr = append(addr, ipNet.IP.String())
					}
				}
			}
		}
	}
	return
}
