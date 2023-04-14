package http

import (
	"log"
	"net/http"
)

type httpHandler struct {
}

func NewHTTPHandler() *httpHandler {
	return &httpHandler{}
}

func (h *httpHandler) ServeContent(addr, route, dir string) (err error) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle(route, http.StripPrefix(route, fs))
	log.Print("HTTP Listening on ", addr)
	err = http.ListenAndServe(addr, nil)
	return
}
