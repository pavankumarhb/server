package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(rw, "pl, %s", d)
	})

	http.HandleFunc("/user", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(rw, "Welcome to cila %s", d)
	})

	http.ListenAndServe(":8000", nil)
}

