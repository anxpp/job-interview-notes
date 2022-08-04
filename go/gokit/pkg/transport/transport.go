package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

// 在文件的第一部分中，我们将请求和响应映射到它们的 JSON 有效负载。
type GetRequest struct{}

type GetResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

type ValidateRequest struct {
	Date string `json:"date"`
}

type ValidateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type StatusRequest struct{}

type StatusResponse struct {
	Status string `json:"status"`
}

// 在第二部分中，我们将为传入的请求编写「解码器」
func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetRequest
	return req, nil
}

func DecodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req StatusRequest
	return req, nil
}

// 最后，我们有响应输出的编码器
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
