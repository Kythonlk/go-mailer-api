package handler

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subject := r.FormValue("subject")
	body := r.FormValue("body")
	recipient := r.FormValue("recipient")

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

	headers := make([]byte, 0)
	headers = append(headers, []byte(fmt.Sprintf("From: %s\r\n", smtpUser))...)
	headers = append(headers, []byte(fmt.Sprintf("To: %s\r\n", recipient))...)
	headers = append(headers, []byte(fmt.Sprintf("Subject: %s\r\n", subject))...)
	headers = append(headers, []byte("\r\n")...)

	message := append(headers, []byte(body)...)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{recipient}, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Email sent successfully!")
}
