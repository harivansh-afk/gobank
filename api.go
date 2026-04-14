package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handlAccount))

	log.Println("JSON API server running on port:", s.listenAddr)		

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handlAccount(w http.ResponseWriter,r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handlCreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("methods not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter,r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter,r *http.Request) error {
	return nil
}

func (s *APIServer) handlCreateAccount(w http.ResponseWriter,r *http.Request) error {
	return nil
}

func (s *APIServer) handlTransfer(w http.ResponseWriter,r *http.Request) error {
	return nil
}


