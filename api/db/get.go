package handler

import (
	"net/http"
	"strings"

	"github.com/jwx233s/a-service/pkg/db"
	"github.com/jwx233s/a-service/pkg/response"
)

// GET /api/db/get/user                    - 查询全部
// GET /api/db/get/user?id=1               - 按 id 查询
// GET /api/db/get/user?json.user_id=123   - 按 json 字段查询
func Handler(w http.ResponseWriter, r *http.Request) {
	response.SetHeaders(w)
	if r.Method == "OPTIONS" {
		return
	}

	// 从路径提取表名: /api/db/get/user -> user
	table := extractTableName(r.URL.Path, "get")
	if table == "" {
		response.Error(w, "Missing table name", 400)
		return
	}
	if !db.AllowedTables[table] {
		response.Error(w, "Table not allowed", 403)
		return
	}

	query := "select=*"
	filter := db.BuildFilter(r)
	if filter != "" {
		query += "&" + filter
	}

	data, err := db.Select(table, query)
	if err != nil {
		response.Error(w, "Query failed", 500)
		return
	}
	response.JSON(w, data)
}

func extractTableName(path, action string) string {
	// /api/db/get/user -> user
	prefix := "/api/db/" + action + "/"
	if strings.HasPrefix(path, prefix) {
		table := strings.TrimPrefix(path, prefix)
		// 去掉可能的尾部斜杠
		return strings.Split(table, "/")[0]
	}
	return ""
}
