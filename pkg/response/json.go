package response

import (
	"encoding/json"
	"net/http"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`    // 状态码: 200 成功, 其他失败
	Data    interface{} `json:"data"`    // 数据
	Message string      `json:"message"` // 消息
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// JSON 返回成功响应，只提取 json 字段内容
func JSON(w http.ResponseWriter, data []byte) {
	SetHeaders(w)

	// 解析原始数据，提取 json 字段
	var rawData []map[string]interface{}
	json.Unmarshal(data, &rawData)

	// 只保留 json 字段的内容
	var result []interface{}
	for _, item := range rawData {
		if jsonField, ok := item["json"]; ok {
			result = append(result, jsonField)
		}
	}

	resp := Response{
		Code:    200,
		Data:    result,
		Message: "success",
	}
	json.NewEncoder(w).Encode(resp)
}

// Success 返回成功响应
func Success(w http.ResponseWriter, data interface{}) {
	SetHeaders(w)
	resp := Response{
		Code:    200,
		Data:    data,
		Message: "success",
	}
	json.NewEncoder(w).Encode(resp)
}

// Error 返回错误响应
func Error(w http.ResponseWriter, msg string, code int) {
	SetHeaders(w)
	w.WriteHeader(code)
	resp := Response{
		Code:    code,
		Data:    nil,
		Message: msg,
	}
	json.NewEncoder(w).Encode(resp)
}
