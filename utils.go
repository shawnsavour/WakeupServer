package main

// Stolen from https://github.com/sabhiram/go-wol

import (
	"log"
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
	// if regex find true
	if matched, _ := regexp.MatchString(regex, string(out)); matched {
		// get ip from regex
		r, _ := regexp.Compile(regex)
		
		return r.FindStringSubmatch(string(out))[1]
	} else {
		log.Fatal("No ip found for mac address")
	}
	return ""
}

// check if machine is on
func isOnline(ip string) bool {
	log.Println("Checking if ip " + ip + " machine is online")
	// if ping is success return true
	if _, err := exec.Command("ping", "-c", "1", ip).Output(); err == nil {
		return true
	} else {
		return false
	}
}