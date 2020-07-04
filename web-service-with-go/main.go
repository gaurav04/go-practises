package main

import "net/http"

type foohandler struct {
	Message string
}

//handler interface have serverHTTP method

func (f *foohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called"))
}

func main() {
	http.Handle("/foo", &foohandler{Message: "Foo Called"})
	http.HandleFunc("/bar", barhandler)
	http.ListenAndServe(":8081", nil)
}
