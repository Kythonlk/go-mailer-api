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

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", recipient, subject, body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{recipient}, []byte(message))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Email sent successfully!")
}
