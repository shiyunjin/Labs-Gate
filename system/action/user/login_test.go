package user

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"github.com/shiyunjin/Labs-Gate/system/db"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/middlewares"
	"github.com/shiyunjin/Labs-Gate/system/util"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.Init()
	util.JwtInit()

	db.Connect()

	server := gin.New()
	server.Use(middlewares.Connect)

	server.POST("/login", Login)

	jsonStr := []byte(`{"username":"test","password":"test","checkbox":false}`)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("login test error")
	}

	if string(body) != "{\"currentAuthority\":\"guest\",\"status\":401,\"statusText\":\"unauthorized\"}" {
		t.Fatalf("body error: %v", string(body))
	}

}