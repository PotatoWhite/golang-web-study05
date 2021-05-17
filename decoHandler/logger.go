package decoHandler

import (
	"log"
	"net/http"
	"time"
)

// HTTP logging Decorator
func ElapseLogger(writer http.ResponseWriter, request *http.Request, handler http.Handler) {
	start := time.Now()
	log.Print("[LOGGER 1] Started")
	handler.ServeHTTP(writer, request)
	log.Print("[LOGGER 1] Completed", time.Since(start).Milliseconds())
}

// HTTP logging Decorator
func ParamLogger(writer http.ResponseWriter, request *http.Request, handler http.Handler) {
	log.Print("[LOGGER 2] Param : ", request.URL.Query(), "Method ", request.Method)
	handler.ServeHTTP(writer, request)
	log.Print("[LOGGER 2] Completed")
}
