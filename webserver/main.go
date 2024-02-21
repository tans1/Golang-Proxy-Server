package main

import (
	"net/http"
)

func documentationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	http.Redirect(w, r, "https://pkg.go.dev/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", documentationHandler)
	http.ListenAndServe(":8081", nil)
}

