package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func main() {
	logger, _ = zap.NewProduction()
	var b interface{}
	s := "{\"a\":\"b\"}"
	_ = json.Unmarshal([]byte(s), &b)
	logger.Info("test s", zap.Any("string", s))
	logger.Info("test o", zap.Reflect("object", b))

	r := gin.Default()
	r.Use(func(c *gin.Context) {
	})
	r.POST("/user/login", ApiUserLogin)
	_ = r.Run(":8963")
}

func ApiUserLogin(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
	_ = c.BindJSON(&body)
	if body.Name == "" {
		c.JSON(http.StatusOK, Failed(500, "用户名不能为空"))
		return
	}
	c.JSON(http.StatusOK, ServiceUserLogin(body.Name, body.Pass))
}

func ServiceUserLogin(name, pass string) Response {
	// business code
	return Success(map[string]string{"name": name, "pass": pass})
}

type Response struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{
		Code: 0,
		Data: data,
	}
}

func Failed(code uint, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}
