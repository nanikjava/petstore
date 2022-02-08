package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test.rest/api/store"
)

func main() {
	r := mux.NewRouter()
	ConfigureRouter(r)
}

//ConfigureRouter to configure the mux router
func ConfigureRouter(r *mux.Router) {
	p := store.Pet{}
	p.ConfigureRoute(r)
	log.Fatal(http.ListenAndServe(":8100", r))
}
