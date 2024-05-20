package client

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

func ExecuteRemoteCommand(serverAddress, username, password, command string) (string, error) {
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
	defer session.Close()
	fmt.Println(strings.TrimRight(command, "\r\n"))
	output, err := session.CombinedOutput(strings.TrimRight(command, "\r\n"))
	if err != nil {
		return "", fmt.Errorf("failed to run command: %s", err)
	}
	return string(output), nil
}
