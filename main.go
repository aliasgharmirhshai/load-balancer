package main

import (
	"fmt"
	"os"
	"github.com/aliasgharmirhshai/load-balancer/receiver"
)

func main() {
	port := ":" + os.Args[1]

	err := receiver.runServer(port)

	if err != nil {
		fmt.Println("\nError starting server:", err)
		return
	}
}


