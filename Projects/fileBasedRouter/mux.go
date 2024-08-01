package main

import (
	"net/http"
)

type filo struct {
	filelist *[]string
}

func (fl *filo) RenderPage(w http.ResponseWriter, r *http.Request) {

}

func newFRouter() (*http.ServeMux, error) {

	routes, err := getList("pages/")
	if err != nil {
		return nil, err
	}
	fl := &filo{filelist: &routes}

	address := "localhost:8080"
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(route, fl.RenderPage)
	}

	http.ListenAndServe(address, mux)

	return mux, nil
}
