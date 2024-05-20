package main

import (
	C "app/client"
	"fmt"
)

func main() {

	username := "oussama"
	password := "oussama"
	serverAddress := "172.21.163.14"

	output, err := C.ExecuteRemoteCommand(serverAddress, username, password, "ifconfig")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println("Command output:")
	fmt.Println(output)
}
