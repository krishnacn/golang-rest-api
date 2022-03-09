package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var muxDispatcher = mux.NewRouter()

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	log.Println("mux server listening on port", port)
	http.ListenAndServe(port, muxDispatcher)
}
