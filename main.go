package main

import (
	"log"
	"net/http"
	"os"
)

var (
	port = os.Getenv("PORT")
)

// cors adds required headers to responses such that direct access works.
//
// These are not required if using "proxy" access.
func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "accept, content-type")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}

func main() {

	if port == "" {
		port = "4444"
	}

	srv := &http.Server{Addr: ":" + port}

	s, err := NewSever()
	if err != nil {
		log.Fatal("error creating server", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", cors(s.handle))
	srv.Handler = mux

	log.Printf("Serving on https://0.0.0.0:" + port)
	log.Fatal(srv.ListenAndServe())
}
