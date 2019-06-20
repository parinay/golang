package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Http")
}

func getform(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<form method=\"POST\" action=\"/postform\"> <input name=\"myname\"/> <input type = \"submit\"/> </form>")
}
func postform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "your name is %s \n", r.PostFormValue("myname"))
}
func someform(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, "<form method=\"POST\" action=\"/postform\"> <input name=\"myname\"/> <input type = \"submit\"/> </form>")
	} else {
		r.ParseForm()
		fmt.Fprintf(w, "your name is %s \n", r.PostFormValue("myname"))
	}
}
func main() {
	log.Println("starting server")
	http.HandleFunc("/company", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "My address") })
	http.Handle("/", MyHandler{})
	http.HandleFunc("/form", getform)
	http.HandleFunc("/postform", postform)
	http.HandleFunc("/someform", someform)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8086", nil)
}
