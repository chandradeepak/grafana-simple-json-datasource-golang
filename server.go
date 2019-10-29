package main

import (
	"log"
	"net/http"
)

type server struct {
}

// NewSever ...
func NewSever() (*server, error) {
	s := &server{}
	return s, nil
}

func (s *server) handle(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Println("got connection")

	path := r.URL.Path
	log.Println("the path is", path)

	switch path {
	case "/":
		log.Println("this is a dummy server for grafana simple json data source")
		// w.Write([]byte("this is a dummy server for grafana simple json data source"))
		w.WriteHeader(http.StatusOK)
		break
	case "/search":
		log.Println("used by the find metric options on the query tab in panels")
		s.search(w, r)
	case "/query":
		log.Println("should return metrics based on input")
		s.query(w, r)
	case "/annotations":
		log.Println("should return annotations.")
		s.annotations(w, r)
	case "tags-keys":
		log.Println("should return tag keys for ad hoc filters.")
		s.tagsKeys(w, r)
	case "tags-values":
		log.Println("should return tag values for ad hoc filters.")
		s.tagsValues(w, r)
	}

}

// These are not required if using "proxy" access.
func (s *server) cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "accept, content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "http://grafana.playland-01-cluster.k8s.cs.swiftnav.com")
}

func (s *server) search(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`["upper_25","upper_50","upper_75","upper_90","upper_95"]`))

}

func (s *server) query(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL)
	log.Println(r.Body)

}

func (s *server) annotations(w http.ResponseWriter, r *http.Request) {

}

func (s *server) tagsKeys(w http.ResponseWriter, r *http.Request) {

}

func (s *server) tagsValues(w http.ResponseWriter, r *http.Request) {

}
