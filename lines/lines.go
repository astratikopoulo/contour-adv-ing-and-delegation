// This code exists only to facilitate a simple demo.
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var photo []byte

const url = "https://images.kusi.com/wp-content/uploads/2018/01/dmv.jpg"

func init() {
	log.Println("Started init")

	// don't worry about errors
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	photo, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Starting lines application.")
	http.HandleFunc("/", mountainHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mountainHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(photo)
	if err != nil {
		log.Printf("An error ocurred returning the image. err: %s", err.Error())
	}
}
