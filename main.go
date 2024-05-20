package main

import (
	C "app/client"
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	fmt.Println("**REMOTE EEXCUTE BASH COMMAND VIA SSH PROTOCOL**")
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan string)

	go func() {
		for {
			fmt.Println("Enter bash command type:")
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("read line: %s \n", line)
			c <- line
		}
	}()

	go func() {
		for {
			fmt.Println("ssh client waiting command to connect")
			command := <-c
			if command == "test" {
				wg.Done()
			}
			username := "oussama"
			password := "oussama"
			serverAddress := "172.18.129.113"
			output, err := C.ExecuteRemoteCommand(serverAddress, username, password, command)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}

			fmt.Println("Command output:")
			fmt.Println(output)
		}

	}()
	wg.Wait()
}
