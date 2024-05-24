package client

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

func ExecuteRemoteCommand(serverAddress string, username string, password string, channel chan string) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", serverAddress), config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	stdinBuf, _ := session.StdinPipe()

	session.Shell()
	srdoutBuf, _ := session.StdoutPipe()
	defer session.Close()
	for {
		command := <-channel
		formattedCommand := strings.TrimRight(command, "\r\n")
		if formattedCommand == "END" {
			fmt.Println("end connection to the server")
			return
		}
		fmt.Println(formattedCommand)
		stdinBuf.Write([]byte(formattedCommand))
		// output, err := session.CombinedOutput(formattedCommand)
		// if err != nil {
		// 	fmt.Printf("failed to run command: %s", err)
		// }
		results := make([]byte, 4000)
		srdoutBuf.Read(results)
		fmt.Println(string(results))
	}
}
