package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	Platform  string `json:"platform"`
	Language  string `json:"language"`
	Timestamp int64  `json:"timestamp"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message:   "Hello World!",
		Status:    "success",
		Platform:  "Vercel",
		Language:  "Go",
		Timestamp: time.Now().UnixMilli(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(resp)
}
