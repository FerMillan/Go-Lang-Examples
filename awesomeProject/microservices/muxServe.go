package microservices

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func ArticlesCategoryHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Call ArticlesCategory Handler")
	vars := mux.Vars(r)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Category: %v\n\n", vars["category"])
	d, err := ioutil.ReadAll(r.Body)
	uri := r.RequestURI
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "URL %v%v\n\n", r.Host, r.URL)
	fmt.Fprintf(rw, "Body: %s\n\n", d)
	fmt.Fprintf(rw, "Params: %v\n\n", r.URL.Query())
	fmt.Fprintf(rw, "URI: %v\n\n", uri)
	fmt.Fprintf(rw, "Headers: %v\n\n", r.Header)
	fmt.Fprintf(rw, "Cookies: %v\n\n", r.Header["Cookie"])
	h := len(r.Header["Cookie"])
	if h > 0 {
		fmt.Fprintf(rw, "SESSID: %v\n\n", strings.Split(r.Header["Cookie"][0], "; ")[0])
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	log.Println("Call Home Handler")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home Handler")
}

func RunConnection() {

	log.Println("Begin Connection")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).
		Methods("GET")
	r.HandleFunc("/articles/{category}", ArticlesCategoryHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:9091",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}() // make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill) // This blocks the code until the channel receives some message
	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)
	// Once message is consumed shut everything down
	// Gracefully shuts down all client requests. Makes server more reliable
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(tc)

	log.Fatal(srv.ListenAndServe())
}
