package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	"github.com/Jeffail/tunny"
)

type fi struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	numCPUs := runtime.NumCPU()

	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {

		// TODO: Something CPU heavy with payload
		bPayLoad, err := GetBytes(payload)
		if err != nil {
			return nil
		}
		gne := &fi{}
		json.Unmarshal(bPayLoad, gne)
		gne.Id = 100
		gne.Name = "helloWorld"

		b, err := json.Marshal(gne)
		if err != nil {
			return nil
		}

		return b
	})
	defer pool.Close()

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		// Funnel this work into our pool. This call is synchronous and will
		// block until the job is completed.
		result := pool.Process(input)

		w.Write(result.([]byte))
	})
	fmt.Println("start listen 8090...")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Panic(err)
	}
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
