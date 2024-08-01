package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIerror struct {
	error string `json: "error"`
}

type APIServer struct {
	listenAddress string
	store         Storage
}

type APIfunc func(http.ResponseWriter, *http.Request) error

func makeHTTPhandler(f APIfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println(err)
			// writeJSON(w, http.StatusBadRequest, APIerror{error: err.Error()})
		}

	}
}

func newAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{listenAddress: listenAddress, store: store}

}

func (s *APIServer) run() {
	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPhandler(s.handleGetAccount)).Methods("GET")
	router.HandleFunc("/account", makeHTTPhandler(s.handleCreateAccount)).Methods("POST")

	fmt.Println("The server has started at port: ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Method not allowed: %s", r.Method)
}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	if id == "" {
		fmt.Println("Missing id")
		return writeJSON(w, http.StatusBadRequest, APIerror{error: "Missing id"})
	}
	accountID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	account, err := s.store.GetAccountByID(accountID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, account)
}

type NewAccountRequest struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	accountRequest := &NewAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(accountRequest); err != nil {
		return err
	}
	id, err := s.store.CreateAccount(accountRequest.FirstName, accountRequest.LastName)
	if err != nil {
		return err
	}
	log.Println("New account created with id: ", id)
	account := newAccount(accountRequest.FirstName, accountRequest.LastName)
	account.ID = id
	return writeJSON(w, http.StatusOK, account)

}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
