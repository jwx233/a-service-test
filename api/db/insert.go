package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/jwx233s/a-service/pkg/db"
	"github.com/jwx233s/a-service/pkg/response"
)

// POST /api/db/insert/user
// Body: {"json": {"user_id": "123", "name": "Tom"}}
func Handler(w http.ResponseWriter, r *http.Request) {
	response.SetHeaders(w)
	if r.Method == "OPTIONS" {
		return
	}

	table := extractTableName(r.URL.Path, "insert")
	if table == "" {
		response.Error(w, "Missing table name", 400)
		return
	}
	if !db.AllowedTables[table] {
		response.Error(w, "Table not allowed", 403)
		return
	}

	body, _ := io.ReadAll(r.Body)
	if len(body) == 0 {
		response.Error(w, "Missing body", 400)
		return
	}

	data, err := db.Insert(table, string(body))
	if err != nil {
		response.Error(w, "Insert failed", 500)
		return
	}
	response.JSON(w, data)
}

func extractTableName(path, action string) string {
	prefix := "/api/db/" + action + "/"
	if strings.HasPrefix(path, prefix) {
		table := strings.TrimPrefix(path, prefix)
		return strings.Split(table, "/")[0]
	}
	return ""
}
