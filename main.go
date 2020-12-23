package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Ooops"))
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":8080", nil)
}
