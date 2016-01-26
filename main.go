package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    s := router.Queries("key", "{key:[a-b][0-9]+}").Subrouter()
    s.HandleFunc("/bullshit", Index)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", r.URL.Query()["key"])
}