package main

import "net/http"

type jsonResponse struct {
	Error bool
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

}
