package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp, _ := http.Get("http://ifconfig.me/")
	defer resp.Body.Close()

	byteArray, _ := io.ReadAll(resp.Body)
	fmt.Fprint(w, string(byteArray))
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
