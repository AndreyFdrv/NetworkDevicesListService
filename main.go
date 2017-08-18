package main

import (
	"fmt"
	"net"
	"net/http"
)

func writeNetworkDevicesList(w http.ResponseWriter, r *http.Request) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "List of network devices:\n\n")
	for i := 0; i < len(interfaces); i++ {
		fmt.Fprintf(w, "Name: %v; ", interfaces[i].Name)
		var address string
		if fmt.Sprintf("%v", interfaces[i].HardwareAddr) == "" {
			address = "-"
		} else {
			address = fmt.Sprintf("%v", interfaces[i].HardwareAddr)
		}
		fmt.Fprintf(w, "Address: %v; ", address)
		fmt.Fprintf(w, "Index: %v; ", interfaces[i].Index)
		fmt.Fprintf(w, "Maximum transmission unit: %v bytes; ", interfaces[i].MTU)
		fmt.Fprintf(w, "Flags: %v\n", interfaces[i].Flags)
	}
}
func main() {
	fmt.Println("Listening on port :3000")
	http.HandleFunc("/", writeNetworkDevicesList)
	http.ListenAndServe(":3000", nil)
}
