package main

import (
	C "app/client"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("**REMOTE EEXCUTE BASH COMMAND VIA SSH PROTOCOL**")
	fmt.Println("Init websocket server")
	//S.InitConnection()
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan string)

	//read ip
	fmt.Println("please enter server address")
	reader := bufio.NewReader(os.Stdin)
	serverAddress, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//read username
	fmt.Println("please enter username")
	reader2 := bufio.NewReader(os.Stdin)
	username, err := reader2.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//raed password
	fmt.Println("please enter password")
	reader3 := bufio.NewReader(os.Stdin)
	password, err := reader3.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

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
		fmt.Println("ssh client start connection to the server")
		C.ExecuteRemoteCommand(strings.TrimRight(serverAddress, "\r\n"), strings.TrimRight(username, "\r\n"), strings.TrimRight(password, "\r\n"), c)
	}()
	wg.Wait()
}
