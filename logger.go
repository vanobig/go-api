package main

import (
	"log"
	"net/http"
	"time"
	"net/url"
)

func Logger(inner Handle) Handle {
	return Handle(func(w http.ResponseWriter, r *http.Request, params url.Values) {
		start := time.Now()

		inner(w, r, params)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}