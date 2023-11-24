package main

import (
	"fmt"
	"os"
	"github.com/roopeshsn/log-ingestor/query-interface/client"
)

func main()  {
	command := os.Args[1]
	switch command {
	case "status":
		client.Status(os.Args[2:])
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}