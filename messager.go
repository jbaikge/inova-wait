package main

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendMessage(body string) (err error) {
	var from, pass string
	var to []string

	if from = os.Getenv("EMAIL_FROM"); from == "" {
		return errors.New("EMAIL_FROM not defined")
	}
	if pass = os.Getenv("EMAIL_PASS"); pass == "" {
		return errors.New("EMAIL_PASS not defined")
	}
	if to = strings.Fields(os.Getenv("EMAIL_TO")); len(to) == 0 {
		return errors.New("EMAIL_TO not defined")
	}

	host := "smtp.gmail.com"
	port := 587
	auth := smtp.PlainAuth("", from, pass, host)

	message := []byte(body)
	address := fmt.Sprintf("%s:%d", host, port)

	return smtp.SendMail(address, auth, from, to, message)
}
