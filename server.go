package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

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
		m, err := jSONIntMap(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		o.Set(m)

	}
}

func main() {
	output := &output.ENTTECUSBProOutput{COMPort: "/dev/tty.usbserial-EN158833"}
	http.HandleFunc("/", makeHandler(output))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
