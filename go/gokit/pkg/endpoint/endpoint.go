package endpoint

import (
	"GoNotes/go/gokit/pkg/service"
	"GoNotes/go/gokit/pkg/transport"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// 公开端点
type Endpoints struct {
	GetEndpoint      endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

// MakeGetEndpoint 返回 「get」服务的响应
func MakeGetEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transport.GetRequest) // 我们只需要请求，不需要使用它的值
		d, err := srv.Get(ctx)
		if err != nil {
			return transport.GetResponse{Date: d, Err: err.Error()}, nil
		}
		return transport.GetResponse{Date: d}, nil
	}
}

// MakeStatusEndpoint 返回 「status」服务的响应
func MakeStatusEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(transport.StatusRequest) // 我们只需要请求，不需要使用它的值
		s, err := srv.Status(ctx)
		if err != nil {
			return transport.StatusResponse{s}, err
		}
		return transport.StatusResponse{s}, nil
	}
}

// MakeValidateEndpoint 返回「validate」服务的响应
func MakeValidateEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.ValidateRequest)
		b, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return transport.ValidateResponse{b, err.Error()}, nil
		}
		return transport.ValidateResponse{b, ""}, nil
	}
}

// get 端点映射
func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := transport.GetRequest{}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(transport.GetResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}

// Status 端点映射
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := transport.StatusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(transport.StatusResponse)
	return statusResp.Status, nil
}

// Validate 端点映射
func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
	req := transport.ValidateRequest{Date: date}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(transport.ValidateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
