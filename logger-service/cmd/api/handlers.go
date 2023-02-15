package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `jason:"name"`
	Data string `jason:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// json oku

	var requestPayload JSONPayload
	_ = app.readJson(w, r, &requestPayload)

	//data ekle

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJson(w, http.StatusAccepted, resp)
}
