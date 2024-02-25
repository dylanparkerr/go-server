package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func (srv *Server) HandleUsersCreate(writer http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodPost, http.MethodPut:
		// check content type
		if req.Header.Get("Content-Type") != "application/json" {
			writer.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		// read json body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("could not read json body: %v", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer req.Body.Close()

		// unmarashal the json into a user sctruct
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			fmt.Printf("could not unmarshal json body: %v", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("User created for: %v\n", user.Name)
		srv.db.AddUser(&user)

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

}
