package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	db *Database
}

func newServer(db *Database) *Server {
	return &Server{
		db: db,
	}
}

func (srv *Server) HandleIndex(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("Custom-Header", "custom-value")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("welcome to the homepage"))
	// probably don't need to try and handle this error here.. but i want practice
	// and writing panic is kind of funny
	if err != nil {
		panic(1)
	}
	fmt.Println("root endpoint hit")
}

func (srv *Server) HandleUsers(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	for _, user := range srv.db.Users {
		jsonUser, jsonErr := json.Marshal(user)
		_, writeErr := writer.Write(jsonUser)

		if jsonErr != nil || writeErr != nil {
			http.Error(writer, "Error creating JSON response", http.StatusInternalServerError)
		}
	}
	fmt.Println("users endpoint hit")
}
