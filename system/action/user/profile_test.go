package user

import (
	"github.com/shiyunjin/Labs-Gate/system/e"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfile(t *testing.T) {
	server := testGinWithLogin()

	server.GET("/profile", Profile)


	req := httptest.NewRequest(http.MethodGet, "/profile", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("get profile test error")
	}

	if string(body) != `{"avatar":"/public/avatar.png","department":"admin","name":"admin"}` {
		t.Fatalf("get profile body error: %v", string(body))
	}
}

func TestProfileNil(t *testing.T) {
	server := testGin()

	server.GET("/profile", Profile)


	req := httptest.NewRequest(http.MethodGet, "/profile", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("get profile nil test error")
	}

	if string(body) != `{"avatar":"/public/avatar.png","department":"guest","name":"???"}` {
		t.Fatalf("get profile body error: %v", string(body))
	}
}

func TestAuthority(t *testing.T) {
	server := testGinWithLogin()

	server.GET("/authority", Authority)


	req := httptest.NewRequest(http.MethodGet, "/authority", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("get profile test error")
	}

	if string(body) != `{"currentAuthority":"admin","status":200,"statusText":"ok"}` {
		t.Fatalf("get profile body error: %v", string(body))
	}
}

func TestAuthorityNil(t *testing.T) {
	server := testGin()

	server.GET("/authority", Authority)


	req := httptest.NewRequest(http.MethodGet, "/authority", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("get profile test error")
	}

	if string(body) != `{"currentAuthority":"guest","status":200,"statusText":"ok"}` {
		t.Fatalf("get profile body error: %v", string(body))
	}
}