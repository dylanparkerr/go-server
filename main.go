package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 'database' setup
	usr := newUser("Dylan", 27)
	db := newDatabase()
	db.AddUser(usr)
	srv := newServer(db)

	fmt.Println("Starting server!")
	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns
	// and calls the handler for the pattern that most closely matches the URL.
	// can use this syntax but i think i will use the verbose syntax until I get the pointer stuff
	// router := http.NewServeMux()
	var router *http.ServeMux = http.NewServeMux()
	// one way to handle routes is to set sturct on the htt.ServeMux.Handle object
	// this struct needs to implement the http.Handler interface by implenting the ServeHttp method

	// this is the other way to add routes, but just implenting the body of what looks like an anonymous function
	// This function is convenient when you want to use a function as a handler without needing to define a full type that implements the http.Handler interface
	router.HandleFunc("/", srv.HandleIndex)
	router.HandleFunc("/users/", srv.HandleUsers)
	router.HandleFunc("/users/create/", srv.HandleUsersCreate)

	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux.
	// Handle and HandleFunc add handlers to DefaultServeMux:
	// http.Handle("/foo", fooHandler)
	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// More control over the server's behavior is available by creating a custom Server:
	var server *http.Server = &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
