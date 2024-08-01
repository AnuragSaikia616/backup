package main

import (
	"flag"
	"log"
)

type Application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	address := flag.String("Port", ":4000", "The port of the server")
	app := &Application{}

	log.Printf("The server has started at port %s", *address)
	e := app.getRoutes()
	if e.Start(*address) != nil {
		log.Fatal(e)
	}

}
