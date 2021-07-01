package main

import(
	"fmt"
	"net"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)){
     	d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(rw, "Welcome to Index Page, %s", d)
	})

	http.HandleFunc("/user", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(rw, "Welcome User %s", d)
	})

	http.ListenAndServe(":8000", nil)

	}
}