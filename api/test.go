package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/jwx233s/a-service/pkg/response"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	response.Success(w,"Hello World")
}
