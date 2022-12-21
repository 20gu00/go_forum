package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 单元测试就是看你的请求响应
func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// engine
	// 这里使用router循环引用
	r := gin.Default()
	r.POST("/api/v1/note", CreatePostHandler)

	// json
	body := `{
		"xx":"xx"
		}`
	// 创建个请求
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/note", bytes.NewReader([]byte(body)))

	// 接收响应
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 跟业务相关,比较是不是预期的响应
	assert.Equal(t, 200, w.Code)
	// containes
	assert.Equal(t, "xx:xx", w.Body.String())
	assert.Equal(t, w.Body.String(), "xx:xx")

	//相应结构体
	//req:=new(ResponseData)
	//反序列化 json->go的结构体
	//json.Umarshal(w.Body.Bytes(),req)
}
