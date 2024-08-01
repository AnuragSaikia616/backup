package main

import (
	"net/http"
)

func getRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/user/{id}", GetUser)
	router.HandleFunc("/users", GetUsers)
	router.HandleFunc("/users/", CreateUser)
	router.HandleFunc("/users/{id}", UpdateUser)
	router.HandleFunc("/users/{id}", DeleteUser)

	return router
}
