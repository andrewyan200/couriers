package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Route construction function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Setting router to serve static files
	// Declare static file directory and point it to /assets folder
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/courier", getCourierHandler).Methods("GET")
	r.HandleFunc("/courier", createCourierHandler).Methods("POST")
	return r
}

func main() {

	//old way to handle http connections, like http module in Python, very barebones
	//http.HandleFunc("/", handler)ß
	
	// new way to handle http connections using Mux. Does so by calling newRouter() constructor
	r := newRouter()
	http.ListenAndServe(":8080", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
