package handlers

import (
	"fmt"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "URL %v%v\n", r.Host, r.URL)
	return
}
