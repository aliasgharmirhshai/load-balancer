package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func runServer(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Server Start is : http://localhost%s/", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil

}

func killPort(port string) error {
	return nil
}

func main() {
	port := ":" + os.Args[1]

	err := runServer(port)
	if err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			err = killPort(port)
			if err != nil {
				fmt.Println("\nError killing process: ", err)
				return
			}
			err = runServer(port)
			if err != nil {
				fmt.Println("\nError starting server:", err)
				return
			}
		} else {
			fmt.Println("Error starting server:", err)
			return
		}
	}
}

