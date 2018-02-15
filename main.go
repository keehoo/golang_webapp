package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"database/sql"
	"io/ioutil"
	"log"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function

var tpl *template.Template 

func init() {
	tpl = template.Must(template.ParseGlob("./assets/*.html"))
	
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/morris", morris).Methods("GET")
	r.HandleFunc("/db", database).Methods("GET")
	r.HandleFunc("/users", users).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")

	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func users(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("http://localhost:8081/api/users/allUsers")
	if err != nil {
	
	}


	robots, err := ioutil.ReadAll(res.Body)
	tpl.ExecuteTemplate(w, "morris.html", string(robots))

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func database(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/mysql")
if err != nil {
    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
}
defer db.Close()

// Open doesn't open a connection. Validate DSN data:
err = db.Ping()
if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func morris(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "morris.html", 42)
}