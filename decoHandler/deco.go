package decoHandler

import "net/http"

type DecoratorFunc func(writer http.ResponseWriter, request *http.Request, handler http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	handler http.Handler
}

func (self *DecoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	self.fn(writer, request, self.handler)
}

func NewDecoHandler(handler http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn : fn,
		handler : handler,
	}
}