package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter) {
	Respond(w, http.StatusMethodNotAllowed, map[string]interface{}{
		"error": http.StatusText(http.StatusMethodNotAllowed),
	})
}

func Error(w http.ResponseWriter, status int, errmsg string) {
	Respond(w, status, map[string]interface{}{
		"error": errmsg,
	})
}

func Respond(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("[WARN] httputil.Respond: %v\n", err)
	}
}
