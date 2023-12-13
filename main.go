package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var (
	PORT = 3000
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	addr := fmt.Sprintf(":%d", PORT)

	log.Printf("Listening on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
