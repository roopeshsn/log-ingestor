package client

import (
	"fmt"
	"net"
	"os"
)

func Status(args []string)  {
	fmt.Println("Checking status...")
	host := "localhost"
	port := args[1]
	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Error connecting to Kafka server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Kafka is running on %s\n", address)
}