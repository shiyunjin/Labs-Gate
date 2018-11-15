package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/db"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestLogin(t *testing.T) {
	db.Connect();

	gin.SetMode(gin.TestMode)
	server := router.Router()


	jsonStr := []byte(`{"username":"test","password":"test","checkbox":false}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(jsonStr))

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