package service

import (
	"context"
	"time"
)

// Service 服务建模为接口
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct{}

// NewService 创建一个新的服务
func NewService() Service {
	return dateService{}
}

// Status 仅仅告诉我们我们的服务是正常的！
func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get 将会返回今天的日期
func (dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

// Validate 将会检查日期是否为今天的日期
func (dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false, err
	}
	return true, nil
}
