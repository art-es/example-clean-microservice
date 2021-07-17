package httphandler

import (
	"log"
	"net/http"
)

func NewServeMux() (*http.ServeMux, error) {
	pres, err := newPresenter()
	if err != nil {
		log.Printf("[ERROR] httphandler: NewServeMux: newPresenter: %v\n", err)
		return nil, err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/actions/doubling", pres.ActionDoubling)
	return mux, nil
}
