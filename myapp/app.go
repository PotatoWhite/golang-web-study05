package myapp

import (
	"fmt"
	"github.com/potatowhite/web/study05/decoHandler"
	"net/http"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return addDecorators(mux)
}

func addDecorators(handler http.Handler) http.Handler {
	handler = decoHandler.NewDecoHandler(handler, decoHandler.ParamLogger)
	handler = decoHandler.NewDecoHandler(handler, decoHandler.ElapseLogger)

	return handler
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

