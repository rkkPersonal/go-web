package test

import (
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	// 获取一个请求实例
	w := httptest.NewRecorder()
	// 构造请求
	// 参数依次是 请求方法、路由、参数
	req, _ := http.NewRequest("GET", "/ping", nil)
	// 执行
	router.ServeHTTP(w, req)
	// 断言
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
