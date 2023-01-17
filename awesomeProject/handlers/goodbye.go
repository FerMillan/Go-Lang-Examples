package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Goodbye handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye Function which gives reference to Goodbye handler
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Test Get Properties to call")
	d, err := ioutil.ReadAll(r.Body)
	uri := r.RequestURI
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "URL %v%v\n", r.Host, r.URL)
	fmt.Fprintf(rw, "Body: %s\n\n", d)
	fmt.Fprintf(rw, "Params: %v\n\n", r.URL.Query())
	fmt.Fprintf(rw, "URI: %v\n\n", uri)
	fmt.Fprintf(rw, "Headers: %v\n\n", r.Header)
	fmt.Fprintf(rw, "Cookies: %v\n\n", r.Header["Cookie"])
	if len(r.Header["Cookies"]) > 0 {
		fmt.Fprintf(rw, "SESSID: %v\n\n", strings.Split(r.Header["Cookie"][0], "; ")[0])
	}
}
