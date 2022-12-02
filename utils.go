package main

// Stolen from https://github.com/sabhiram/go-wol

import (
	"log"
	"net"
	"regexp"
	"os/exec"
)



//  using os exec command arp to get first ip from mac
func getIpFromMac(mac string) string {
	cmd := exec.Command("arp", "-a")
	// grep  by mac address and return ip
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	regex := "\\(([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})\\).*" + mac
	r, _ := regexp.Compile(regex)
	ip := r.FindStringSubmatch(string(out))[1] 
	return ip
}

// check if machine is on
func isOnline(ip string) bool {
	_, err := net.DialTimeout("tcp", ip+":80", 1)
	if err != nil {
		return false
	}
	return true
}