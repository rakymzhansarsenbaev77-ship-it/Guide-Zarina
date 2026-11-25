package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	files := http.FileServer(http.Dir("."))
	http.Handle("/", files)
	log.Printf("Ciao. Zarina! is running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
