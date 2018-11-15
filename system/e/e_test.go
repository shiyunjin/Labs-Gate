package e

import "testing"

func TestGetMsg(t *testing.T) {
	if GetMsg(SUCCESS) != "ok" {
		t.Fatalf("e code has error")
	}

	if GetMsg(0) != "fail" {
		t.Fatalf("null e code passed")
	}
}