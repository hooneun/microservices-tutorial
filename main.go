package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hooneun/microservices-tutorial/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	h := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", h)

	http.ListenAndServe(":8080", sm)
}
