package main

import (
	"fmt"
	"net/http"
)

func HandleForm(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}

	name := r.FormValue("name")
	Email := r.FormValue("email")
	fmt.Fprintf(w, "Name=%s\n Email=%s\n", name, Email)

}

func main() {

	fmt.Println("Simple Server")
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileServer)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bahi")

	})

	http.HandleFunc("/form", HandleForm)

	http.ListenAndServe(":8080", nil)

}
