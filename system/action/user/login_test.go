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

func testGin() (server *gin.Engine) {
	gin.SetMode(gin.TestMode)

	config.Init()
	util.JwtInit()

	db.Connect()

	server = gin.New()
	server.Use(middlewares.Connect)

	return server
}

func TestLogin(t *testing.T) {
	server := testGin()

	server.POST("/login", Login)

	jsonStr := []byte(`{"username":"admin","password":"admin","checkbox":true}`)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("login test error")
	}

	if string(body) != `{"currentAuthority":"admin","status":200,"statusText":"ok"}` {
		t.Fatalf("login body error: %v", string(body))
	}

}

func TestLoginEPass(t *testing.T) {
	server := testGin()

	server.POST("/login", Login)

	jsonStr := []byte(`{"username":"admin","password":"admin123","checkbox":true}`)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("login err pass test error")
	}

	if string(body) != `{"currentAuthority":"guest","status":401,"statusText":"unauthorized"}` {
		t.Fatalf("login err pass body error: %v", string(body))
	}

}

func TestLoginDummy(t *testing.T) {
	server := testGin()

	server.POST("/login", Login)

	jsonStr := []byte(`{"username":"testnull","password":"testnull","checkbox":false}`)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("login dummy test error")
	}

	if string(body) != `{"currentAuthority":"guest","status":401,"statusText":"unauthorized"}` {
		t.Fatalf("login dummy body error: %v", string(body))
	}

}

func TestLogout(t *testing.T) {
	server := testGin()

	server.POST("/logout", Logout)

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("logout test error")
	}

}