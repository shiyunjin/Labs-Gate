package rom

import (
	"github.com/shiyunjin/Labs-Gate/system/e"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDel(t *testing.T) {
	server := testGinWithLogin()

	server.POST("/:code/machine/:ip/del", Del)

	req := httptest.NewRequest(http.MethodPost, "/dx101/machine/192.168.0.99/del", nil)

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != e.SUCCESS {
		t.Fatalf("change pass test error")
	}
}
