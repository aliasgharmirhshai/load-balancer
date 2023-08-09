package receiver

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		log.Printf("User %s visited the homepage\n", username)
		fmt.Fprintf(w, "Hello, %s!", username)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Server Start is: http://localhost%s/\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil
}