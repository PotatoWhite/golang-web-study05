package main

import (
	"github.com/potatowhite/web/study05/myapp"
	"net/http"
)

func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}

func NewHandler() http.Handler {
	handler := myapp.NewHandler()
	return handler

}
