package main

import (
	"fmt"
	"log"
)

func main() {
	psqlInfo := "host=localhost port=5432 user=postgres password=gobank dbname=postgres sslmode=disable"
	store, err := NewPostgresStore(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err = store.Init(); err != nil {
		log.Fatal(err)
	}
	defer store.DB.Close()
	fmt.Printf("%+v", store)
	server := newAPIServer(":8080", store)
	server.run()

}
