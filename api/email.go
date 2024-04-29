package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Email(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello World from Email! 👋")
	fmt.Fprintf(w, "Date: %s", time.Now().Format(time.RFC850))
	fmt.Fprintf(w, "Github: https://github.com/riccardogiorato/template-go-vercel/blob/main/api/email.go")
	fmt.Fprintf(w, "Language: go")
	fmt.Fprintf(w, "Cloud: Hosted on Vercel! ▲")
}
