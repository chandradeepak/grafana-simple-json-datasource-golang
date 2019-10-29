package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type server struct {
}

// NewSever ...
func NewSever() (*server, error) {
	s := &server{}
	return s, nil
}

type searchResponse []string

type queryRequest struct {
	Range struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"range"`
	Interval string `json:"interval"`
	Targets  []struct {
		RefID  string `json:"refId"`
		Target string `json:"target"`
	} `json:"targets"`
	Format        string `json:"format"`
	MaxDataPoints int    `json:"maxDataPoints"`
}

type queryTimeSeriesResponse []struct {
	Target     string          `json:"target"`
	Datapoints [][]interface{} `json:"datapoints"`
}

// TableResponseColumn ...
type TableResponseColumn struct {
	Text string `json:"text"`
	Type string `json:"type,omitempty"`
	Sort bool   `json:"sort,omitempty"`
	Desc bool   `json:"desc,omitempty"`
}

// QueryTableResponse ...
type QueryTableResponse []struct {
	Columns []TableResponseColumn `json:"columns"`
	Rows    [][]string            `json:"rows"`
	Type    string                `json:"type"`
}

// AnnotationRqt ...
type AnnotationRqt struct {
	Datasource string `json:"datasource"`
	Enable     bool   `json:"enable"`
	Name       string `json:"name"`
}

// AnnotationRequest ...
type AnnotationRequest struct {
	Range struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"range"`
	RangeRaw struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"rangeRaw"`
	Annotation AnnotationRqt `json:"annotation"`
	Dashboard  string        `json:"dashboard"`
}

// AnnotationRsp ...
type AnnotationRsp struct {
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	Datasource string `json:"datasource"`
}

// AnnotationResponse ...
type AnnotationResponse []struct {
	Annotation AnnotationRsp `json:"annotation"`
	Title      string        `json:"title"`
	Time       int64         `json:"time"`
	Text       string        `json:"text"`
	Tags       []string      `json:"tags"`
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

	decoder := json.NewDecoder(r.Body)
	var qR queryRequest
	err := decoder.Decode(&qR)
	if err != nil {
		log.Println("error decoding query", err)
	}
	log.Printf("printing the query request%+v", qR)

	qRsp := make(QueryTableResponse, 1)
	qRsp[0].Columns = append(qRsp[0].Columns, TableResponseColumn{Text: "column1"})
	qRsp[0].Columns = append(qRsp[0].Columns, TableResponseColumn{Text: "column2"})
	qRsp[0].Columns = append(qRsp[0].Columns, TableResponseColumn{Text: "column3"})

	for i := 0; i < 10; i++ {
		dataRows := make([]string, 3)
		for j := range dataRows {
			dataRows[j] = strconv.Itoa(i * j)
		}
		qRsp[0].Rows = append(qRsp[0].Rows, dataRows)
	}
	qRsp[0].Type = "table"

	rsp, _ := json.Marshal(qRsp)
	//	log.Printf("printing the query response%+v", qRsp)

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, rsp, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	log.Println("printing the query response:", string(prettyJSON.Bytes()))
	w.Write(rsp)

	w.WriteHeader(http.StatusOK)

}

func (s *server) annotations(w http.ResponseWriter, r *http.Request) {

}

func (s *server) tagsKeys(w http.ResponseWriter, r *http.Request) {

}

func (s *server) tagsValues(w http.ResponseWriter, r *http.Request) {

}
