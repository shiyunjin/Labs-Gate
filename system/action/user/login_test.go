package user

import (
	"bytes"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


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

func TestLoginNil(t *testing.T) {
	server := testGin()

	server.POST("/login", Login)


	req := httptest.NewRequest(http.MethodPost, "/login", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	_ = w.Result()

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