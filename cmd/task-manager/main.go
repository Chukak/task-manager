package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// todo
	fs := http.FileServer(http.Dir(os.Getenv("CURRENT_SOURCE_PATH") + "web/templates"))
	log.Print(os.Getenv("CURRENT_SOURCE_PATH") + "web/templates")
	http.Handle("/", fs)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
