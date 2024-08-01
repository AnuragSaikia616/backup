package main

import "net/http"

type APIserver struct {
	host string
	port string
}

func newAPIServer(host string, port string) *APIserver {
	return &APIserver{
		host: host,
		port: port,
	}

}

func (s *APIserver) Run() {
	router := getRoutes()
	err := http.ListenAndServe(s.host+":"+s.port, router)
	if err != nil {
		panic(err)
	}
}

func main() {
	server := newAPIServer("localhost", "8080")
	server.Run()

}
