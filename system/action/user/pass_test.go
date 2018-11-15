package user

import (
	"bytes"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPass(t *testing.T) {
	server := testGinWithLogin()

	server.POST("/pass", Pass)

	jsonStr := []byte(`{"oldpassword":"12345678","newpassword":"12345678","repassword":"12345678"}`)

	req := httptest.NewRequest(http.MethodPost, "/pass", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("change pass test error")
	}
}

func TestPassDummy(t *testing.T) {
	server := testGinWithLogin()

	server.POST("/pass", Pass)

	jsonStr := []byte(`{"oldpassword":"87654321","newpassword":"12345678","repassword":"12345678"}`)

	req := httptest.NewRequest(http.MethodPost, "/pass", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("change pass test error")
	}
}