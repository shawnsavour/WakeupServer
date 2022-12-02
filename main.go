package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ComputerList contains all Computers who we can use to work with
var ComputerList []Computer

func main() {
	log.Println("Starting WakeOnLan Server")
	log.Println(getIpFromMac("2c:f0:5d:33:b7:ab"))
	log.Println("=========================")
	// Start Processing Shell Arguments or use Default Values defined in const.go
	httpPort, computerFilePath := processShellArgs()

	// Process Environment Variables
	httpPort, computerFilePath = processEnvVars(httpPort, computerFilePath)

	// Loading Computer CSV File to Memory File in Memory
	var loadComputerJsonFileError error
	// if ComputerList, loadComputerCSVFileError = loadComputerList(computerFilePath); loadComputerCSVFileError != nil {
	if ComputerList, loadComputerJsonFileError = loadComputerListFromJson(computerFilePath); loadComputerJsonFileError != nil {
		log.Fatalf("Error on loading Computerlist File \"%s\" check File access and formating", computerFilePath)
	}

	// Init HTTP Router - mux
	router := mux.NewRouter()

	// Define Home Route
	router.HandleFunc("/", renderHomePage).Methods("GET")

	// Define Wakeup Api functions with a Computer Name
	router.HandleFunc("/api/wakeup/computer/{computerName}", restWakeUpWithComputerName).Methods("GET")
	router.HandleFunc("/api/wakeup/computer/{computerName}/", restWakeUpWithComputerName).Methods("GET")
	// wake up all computers
	router.HandleFunc("/api/wakeup/all", restWakeUpAll).Methods("GET")
	// Add a Computer to the ComputerList
	router.HandleFunc("/api/computer/add", restAddComputer).Methods("POST")
	// Delete a Computer from the ComputerList
	router.HandleFunc("/api/computer/{computerName}", restDeleteComputer).Methods("DELETE")
	// check if computer is online
	// router.HandleFunc("/api/computer/{computerName}/online", restCheckComputerOnline).Methods("GET")
	// router.HandleFunc("/api/check/computer/{computerName}", restCheckComputer).Methods("GET")

	// Setup Webserver
	httpListen := fmt.Sprint(":", httpPort)
	log.Printf("Startup Webserver on \"%s\"", httpListen)

	log.Fatal(http.ListenAndServe(httpListen, handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router)))
}

//  using os exec command arp to get first ip from mac
func getIpFromMac(mac string) string {
	cmd := exec.Command("arp", "-a | grep " + mac)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
	}
	return string(out)
}