package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Mail sent successfully",
	}

	app.writeJson(w, http.StatusAccepted, payload)
}
