package network_extension

import (
	"net"
	"strings"
	"os/exec"
	"time"
	"fmt"
)

func GetLocalIps() []string {
	var ips []string

	address, err := net.InterfaceAddrs()
	if nil != err {
		return ips
	}

	for _, addr := range address {
		if strings.Contains(addr.String(), ":") {
			continue
		}
		ip := strings.Split(addr.String(), "/")[0]
		ips = append(ips, ip)
	}

	return ips
}

func GetMacAddress(nicName string) string{
	mac := ""
	interfaces, _ := net.Interfaces()

	for _, nic := range interfaces {
		if nic.Name == nicName{
			mac = nic.HardwareAddr.String()
			break
		}
	}

	return mac
}

func IsNetWorkOK(destIpOrDomain string) bool {
	cmd := exec.Command("ping", destIpOrDomain, "-c", "1", "-W", "5")
	startMSTime := time.Now().UnixNano()/1e6
	err := cmd.Run()
	endMSTime := time.Now().UnixNano()/1e6
	if err != nil {
		info := fmt.Sprintf("ping %s failed", destIpOrDomain)
		fmt.Println(info)
		return false
	} else {
		info := fmt.Sprintf("network check time cost %d ms, network status true", endMSTime - startMSTime)
		fmt.Println(info)
	}
	return true
}

func BlockToWaitNetworkReady(destIpOrDomain string){
	isNicReady := false
	for{
		isNicReady = IsNetWorkOK(destIpOrDomain)
		if isNicReady{
			break
		}
		time.Sleep(3 * time.Second)
	}
}