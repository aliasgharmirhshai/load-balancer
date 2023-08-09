package main

import (
	"fmt"
	"net/http"
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


