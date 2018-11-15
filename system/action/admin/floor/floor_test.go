package floor

import (
	"bytes"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {
	server := testGin()

	server.POST("/add", Add)

	jsonStr := []byte(`{"name":"test"}`)

	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("add floor test error")
	}
}


func TestList(t *testing.T) {
	server := testGin()

	server.GET("/floor", List)

	req := httptest.NewRequest(http.MethodGet, "/floor", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("list floor test error")
	}
}

func TestEdit(t *testing.T) {
	server := testGin()

	server.POST("/edit", Edit)

	jsonStr := []byte(`{"id":"5be5960bcc3486de88e69221","name":"test"}`)

	req := httptest.NewRequest(http.MethodPost, "/edit", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("edit floor test error")
	}
}

func TestDel(t *testing.T) {
	server := testGin()

	server.POST("/del", Del)

	jsonStr := []byte(`{"id":"5be5960bcc3486de88e69221"}`)

	req := httptest.NewRequest(http.MethodPost, "/del", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("del floor test error")
	}
}

func TestAddNil(t *testing.T) {
	server := testGin()

	server.POST("/add", Add)

	req := httptest.NewRequest(http.MethodPost, "/add", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	_ = w.Result()
}

func TestEditNil(t *testing.T) {
	server := testGin()

	server.POST("/edit", Edit)

	req := httptest.NewRequest(http.MethodPost, "/edit", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	_ = w.Result()
}

func TestDelNil(t *testing.T) {
	server := testGin()

	server.POST("/del", Del)

	req := httptest.NewRequest(http.MethodPost, "/del", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	_ = w.Result()
}
