package db

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ============ 配置（修改这里） ============

const (
	SUPABASE_URL = "https://arwnlqnzofqqxjnlvgqm.supabase.co"
	SUPABASE_KEY = "sb_publishable_eGii7SaCtXirDh2O5suItQ_L_eZtxuJ"
)

// 允许操作的表
var AllowedTables = map[string]bool{
	"user":      true,
	"community": true,
}

// ============ CRUD 方法 ============

func Select(table, query string) ([]byte, error) {
	endpoint := SUPABASE_URL + "/rest/v1/" + table + "?" + query
	return request("GET", endpoint, "")
}

func Insert(table, body string) ([]byte, error) {
	endpoint := SUPABASE_URL + "/rest/v1/" + table
	return request("POST", endpoint, body)
}

func Update(table, filter, body string) ([]byte, error) {
	endpoint := SUPABASE_URL + "/rest/v1/" + table + "?" + filter
	return request("PATCH", endpoint, body)
}

func Delete(table, filter string) ([]byte, error) {
	endpoint := SUPABASE_URL + "/rest/v1/" + table + "?" + filter
	return request("DELETE", endpoint, "")
}

// 构建过滤条件（支持 id 和 jsonb 字段搜索）
// 示例:
//   ?id=1                     -> id=eq.1
//   ?json.user_id=123         -> json->user_id=eq.123
//   ?json.name=Tom            -> json->name=eq.Tom
//   ?json.status=cs.active    -> json->status=cs.active (contains)
func BuildFilter(r *http.Request) string {
	var filters []string

	for key, values := range r.URL.Query() {
		if len(values) == 0 {
			continue
		}
		value := values[0]

		if key == "id" {
			// 普通 id 查询
			filters = append(filters, "id=eq."+value)
		} else if strings.HasPrefix(key, "json.") {
			// jsonb 字段查询: json.user_id=123 -> json->>user_id=eq.123
			jsonKey := strings.TrimPrefix(key, "json.")
			// 检查是否有操作符前缀 (eq, neq, gt, lt, gte, lte, like, cs, cd)
			if strings.Contains(value, ".") {
				// 已包含操作符: json.status=cs.active
				filters = append(filters, "json->>"+jsonKey+"="+value)
			} else {
				// 默认 eq 操作符
				filters = append(filters, "json->>"+jsonKey+"=eq."+url.QueryEscape(value))
			}
		}
	}

	return strings.Join(filters, "&")
}

// ============ 内部方法 ============

func request(method, endpoint, body string) ([]byte, error) {
	var reqBody io.Reader
	if body != "" {
		reqBody = strings.NewReader(body)
	}

	req, _ := http.NewRequest(method, endpoint, reqBody)
	req.Header.Set("apikey", SUPABASE_KEY)
	req.Header.Set("Authorization", "Bearer "+SUPABASE_KEY)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
