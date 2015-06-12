package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/saulshanabrook/lige/output"
)

// Takes json that is a flat map of int -> int and turns it into a Go map
func jSONIntMap(b io.ReadCloser) (m map[int]int, err error) {
	decoder := json.NewDecoder(b)
	var t map[string]string
	err = decoder.Decode(&t)
	if err != nil {
		return
	}
	m = make(map[int]int)
	for kT, vT := range t {
		var k, v int
		k, err = strconv.Atoi(kT)
		if err != nil {
			return
		}
		v, err = strconv.Atoi(vT)
		if err != nil {
			return
		}
		m[k] = v
	}
	return
}

func makeHandler(o output.Output) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// set the `Set` method on the pass in output, with the json
		// of the request turned into a map
		m, err := jSONIntMap(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		o.Set(m)

	}
}

func main() {
	log.WithFields(log.Fields{
		"port":    "8080",
		"COMPort": os.Args[1],
	}).Info("Starting HTTP Server")
	output := &output.ENTTECUSBProOutput{COMPort: os.Args[1]}
	http.HandleFunc("/", makeHandler(output))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
