// This code exists only to facilitate a simple demo.
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting ping application.")
	http.HandleFunc("/", mountainHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mountainHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong\n"))
	if err != nil {
		log.Printf("An error ocurred returning the image. err: %s", err.Error())
	}
}
