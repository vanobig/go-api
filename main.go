package main

import (
	"net/http"
	"log"
)

func main() {
	r := NewRouter(Root)
	log.Fatal(http.ListenAndServe(":9090", r))
}
