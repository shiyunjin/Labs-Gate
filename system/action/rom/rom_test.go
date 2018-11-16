package rom

import (
	"github.com/shiyunjin/Labs-Gate/system/e"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	server := testGinWithLogin()

	server.GET("/roms", List)


	req := httptest.NewRequest(http.MethodGet, "/roms", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("get rom list test error")
	}
}

func TestMachine(t *testing.T) {
	server := testGinWithLogin()

	server.GET("/:code/machine", Machine)


	req := httptest.NewRequest(http.MethodGet, "/dx603/machine", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("get rom list test error")
	}
}
