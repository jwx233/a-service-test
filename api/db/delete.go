package handler

import (
	"net/http"
	"strings"

	"github.com/jwx233s/a-service/pkg/db"
	"github.com/jwx233s/a-service/pkg/response"
)

// POST /api/db/delete/user?id=1
// POST /api/db/delete/user?json.user_id=123
func Handler(w http.ResponseWriter, r *http.Request) {
	response.SetHeaders(w)
	if r.Method == "OPTIONS" {
		return
	}

	table := extractTableName(r.URL.Path, "delete")
	if table == "" {
		response.Error(w, "Missing table name", 400)
		return
	}
	if !db.AllowedTables[table] {
		response.Error(w, "Table not allowed", 403)
		return
	}

	filter := db.BuildFilter(r)
	if filter == "" {
		response.Error(w, "Missing filter (id or json.xxx)", 400)
		return
	}

	data, err := db.Delete(table, filter)
	if err != nil {
		response.Error(w, "Delete failed", 500)
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
