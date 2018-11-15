package e

import "testing"

func TestGetMsg(t *testing.T) {
	if GetMsg(SUCCESS) != "ok" {
		t.Fatalf("e code has error")
	}
}