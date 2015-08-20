package keydb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// servehttp START OMIT
func (d *DB) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	switch req.Method {
	case "GET":
		d.HandleGet(key, w, req)
	case "POST":
		d.HandleSet(key, w, req)
	default:
		http.Error(w, "Unsupported method", http.StatusBadRequest)
	}
}

// servehttp END OMIT

// get START OMIT
func (d *DB) HandleGet(key string, w http.ResponseWriter, req *http.Request) {
	result, ok := d.Get(key)
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	enc := json.NewEncoder(w)
	enc.Encode(result)
}

// get END OMIT

// set START OMIT
func (d *DB) HandleSet(key string, w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newValue := map[string]interface{}{}
	err = json.Unmarshal(body, &newValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exists := d.Set(key, newValue)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, "%t", exists)
}

// set END OMIT
