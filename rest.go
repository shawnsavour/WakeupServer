// Rest API Implementations

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

//restWakeUpWithComputerName - REST Handler for Processing URLS /api/computer/<computerName>
func restWakeUpWithComputerName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	computerName := vars["computerName"]

	var result WakeUpResponseObject
	result.Success = false

	// Ensure computerName is not empty
	if computerName == "" {
		result.Message = "Empty Computername is not allowed"
		result.ErrorObject = nil
		w.WriteHeader(http.StatusBadRequest)
		// Computername is empty
	} else {

		// Get Computer from List
		for _, c := range ComputerList {
			if c.Name == computerName {

				// We found the Computername
				if err := SendMagicPacket(c.Mac, c.BroadcastIPAddress, ""); err != nil {
					// We got an internal Error on SendMagicPacket
					w.WriteHeader(http.StatusInternalServerError)
					result.Success = false
					result.Message = "Internal error on Sending the Magic Packet"
					result.ErrorObject = err
				} else {
					// Horray we send the WOL Packet succesfully
					result.Success = true
					result.Message = fmt.Sprintf("Succesfully Wakeup Computer %s with Mac %s on Broadcast IP %s", c.Name, c.Mac, c.BroadcastIPAddress)
					result.ErrorObject = nil
				}
			}
		}

		if result.Success == false && result.ErrorObject == nil {
			// We could not find the Computername
			w.WriteHeader(http.StatusNotFound)
			result.Message = fmt.Sprintf("Computername %s could not be found", computerName)
		}
	}
	json.NewEncoder(w).Encode(result)
}

// add a new computer to the list
func restAddComputer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var result WakeUpResponseObject
	result.Success = false

	// Ensure computerName is not empty
	if r.Body == nil {
		result.Message = "Empty Body is not allowed"
		result.ErrorObject = nil
		w.WriteHeader(http.StatusBadRequest)
		// Computername is empty
	} else {

		// Get Computer from List
		var newComputer Computer
		// print request body to log
		fmt.Println(r)

		err := json.NewDecoder(r.Body).Decode(&newComputer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			result.Success = false
			// result.Message = "Error on decoding JSON Body"
			// read r.Body to string
			result.Message = fmt.Sprintf("Error on decoding JSON Body: %s", r.Body)
			result.ErrorObject = err
		} else {
			// Horray we send the WOL Packet succesfully
			result.Success = true
			result.Message = fmt.Sprintf("Succesfully added Computer %s with Mac %s on Broadcast IP %s", newComputer.Name, newComputer.Mac, newComputer.BroadcastIPAddress)
			result.ErrorObject = nil
			ComputerList = append(ComputerList, newComputer)
			// save the new list
			saveComputerList()
		}

	}
	json.NewEncoder(w).Encode(result)
}

// delete a computer from the list by name
func restDeleteComputer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var result WakeUpResponseObject
	result.Success = false

	// Get computer name from URL
	vars := mux.Vars(r)
	computerName := vars["computerName"]

	// Ensure computerName is not empty
	if computerName == "" {

		result.Message = "Empty Computername is not allowed"
		result.ErrorObject = nil
		w.WriteHeader(http.StatusBadRequest)

	} else {

		// Get Computer from List
		for i, c := range ComputerList {
			if c.Name == computerName {

				// We found the Computername
				// remove the computer from the list
				ComputerList = append(ComputerList[:i], ComputerList[i+1:]...)
				// save the new list
				saveComputerList()
				// Horray we send the WOL Packet succesfully
				result.Success = true
				result.Message = fmt.Sprintf("Succesfully deleted Computer %s with Mac %s on Broadcast IP %s", c.Name, c.Mac, c.BroadcastIPAddress)
				result.ErrorObject = nil
			}
		}

		if result.Success == false && result.ErrorObject == nil {
			// We could not find the Computername
			w.WriteHeader(http.StatusNotFound)
			result.Message = fmt.Sprintf("Computername %s could not be found", computerName)
		}
	}


	json.NewEncoder(w).Encode(result)
}





// wake up all computers
func restWakeUpAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var result WakeUpResponseObject
	result.Success = false

	for _, c := range ComputerList {

		// We found the Computername
		if err := SendMagicPacket(c.Mac, c.BroadcastIPAddress, ""); err != nil {
			// We got an internal Error on SendMagicPacket
			w.WriteHeader(http.StatusInternalServerError)
			result.Success = false
			result.Message = "Internal error on Sending the Magic Packet"
			result.ErrorObject = err
		} else {
			// Horray we send the WOL Packet succesfully
			result.Success = true
			result.Message = fmt.Sprintf("Succesfully Wakeup Computer %s with Mac %s on Broadcast IP %s", c.Name, c.Mac, c.BroadcastIPAddress)
			result.ErrorObject = nil
		}
	}

	json.NewEncoder(w).Encode(result)
}

// saveComputerList - Save the ComputerList to the File
func saveComputerList() {
	// Save the ComputerList /app/computerlist.json
	file, _ := json.MarshalIndent(ComputerList, "", " ")
	_ = ioutil.WriteFile("computerlist.json", file, 0644)

}

// check if computer is online
// func restCheckComputerOnline(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	var result WakeUpResponseObject
// 	result.Success = false

// 	// Get computer name from URL
// 	vars := mux.Vars(r)
// 	computerName := vars["computerName"]

// 	// Ensure computerName is not empty
// 	if computerName == "" {

// 		result.Message = "Empty Computername is not allowed"
// 		result.ErrorObject = nil
// 		w.WriteHeader(http.StatusBadRequest)

// 	} else {

// 		// Get Computer from List
// 		for _, c := range ComputerList {
// 			if c.Name == computerName {

// 				// We found the Computername
// 				// check if computer is online
// 				if isOnline(c.IPAddress) {
// 					// Horray we send the WOL Packet succesfully
// 					result.Success = true
// 					result.Message = fmt.Sprintf("Computer %s with IP %s is online", c.Name, c.IPAddress)
// 					result.ErrorObject = nil
// 				} else {
// 					// Computer is not online
// 					result.Success = false
// 					result.Message = fmt.Sprintf("Computer %s with IP %s is offline", c.Name, c.IPAddress)
// 					result.ErrorObject = nil
// 				}
// 			}
// 		}

// 		if result.Success == false && result.ErrorObject == nil {
// 			// We could not find the Computername
// 			w.WriteHeader(http.StatusNotFound)
// 			result.Message = fmt.Sprintf("Computername %s could not be found", computerName)
// 		}
// 	}
// 	json.NewEncoder(w).Encode(result)
// }

// func isOnline(ip string) bool {
// 	_, err := net.DialTimeout("tcp", ip+":22", 1*time.Second)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }