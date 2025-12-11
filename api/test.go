package handler

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// GET /api/test - 查询 test 表所有数据
func Handler(w http.ResponseWriter, r *http.Request) {
	data, err := supabaseSelect("test", "select=*")
	if err != nil {
		jsonError(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	jsonRaw(w, data)
}

// ============ Supabase 工具 ============

func getConfig() (string, string) {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")
	if url == "" {
		url = "https://arwnlqnzofqqxjnlvgqm.supabase.co"
	}
	if key == "" {
		key = "sb_publishable_eGii7SaCtXirDh2O5suItQ_L_eZtxuJ"
	}
	return url, key
}

func supabaseRequest(method, endpoint, body string) ([]byte, error) {
	url, key := getConfig()
	var reqBody io.Reader
	if body != "" {
		reqBody = strings.NewReader(body)
	}

	req, _ := http.NewRequest(method, url+endpoint, reqBody)
	req.Header.Set("apikey", key)
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// 查询: supabaseSelect("users", "select=*&id=eq.1")
func supabaseSelect(table, query string) ([]byte, error) {
	return supabaseRequest("GET", "/rest/v1/"+table+"?"+query, "")
}

// 新增: supabaseInsert("users", `{"name":"Tom"}`)
func supabaseInsert(table, jsonData string) ([]byte, error) {
	return supabaseRequest("POST", "/rest/v1/"+table, jsonData)
}

// 更新: supabaseUpdate("users", "id=eq.1", `{"name":"Jerry"}`)
func supabaseUpdate(table, filter, jsonData string) ([]byte, error) {
	return supabaseRequest("PATCH", "/rest/v1/"+table+"?"+filter, jsonData)
}

// 删除: supabaseDelete("users", "id=eq.1")
func supabaseDelete(table, filter string) ([]byte, error) {
	return supabaseRequest("DELETE", "/rest/v1/"+table+"?"+filter, "")
}

// ============ 响应工具 ============

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func jsonRaw(w http.ResponseWriter, data []byte) {
	setHeaders(w)
	w.Write(data)
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	setHeaders(w)
	w.WriteHeader(code)
	w.Write([]byte(`{"error":"` + msg + `"}`))
}
