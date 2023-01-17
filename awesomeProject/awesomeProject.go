package main

import (
	"awesomeProject/modules/api_calls"
	"fmt"
)

func main() {

	fmt.Println(api_calls.Post())
	//l := log.New(os.Stdout, "Microservice Test ", log.LstdFlags)
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	//homehandler := handlers.NewTestHandler()
	//
	//// There are 2 microservices
	//sm := http.NewServeMux()
	//sm.Handle("/hello", hh)
	//sm.Handle("/goodbye", gh)
	//sm.Handle("/", homehandler)
	//
	//// Config to run the server
	//s := &http.Server{
	//	Addr:         ":9090",
	//	Handler:      sm,
	//	IdleTimeout:  120 * time.Second,
	//	ReadTimeout:  1 * time.Second,
	//	WriteTimeout: 1 * time.Second,
	//}
	//// wrapping ListenAndServe in gofunc so it's not going to block
	//go func() {
	//	err := s.ListenAndServe()
	//	if err != nil {
	//		l.Fatal(err)
	//	}
	//}() // make a new channel to notify on os interrupt of server (ctrl + C)
	//sigChan := make(chan os.Signal)
	//signal.Notify(sigChan, os.Interrupt)
	//signal.Notify(sigChan, os.Kill) // This blocks the code until the channel receives some message
	//sig := <-sigChan
	//l.Println("Received terminate, graceful shutdown", sig)
	//// Once message is consumed shut everything down
	//// Gracefully shuts down all client requests. Makes server more reliable
	//tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()
	//s.Shutdown(tc)
}
