package util

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestIf(t *testing.T) {
	if If(true,"1","2") == "2" {
		t.Fatalf("util if has error")
	}
	if If(false,"1","2") == "1" {
		t.Fatalf("util if has error")
	}
}

func TestIn_array(t *testing.T) {
	if In_array("1",[]string{"2","3"}) {
		t.Fatalf("util In_array has error")
	}
	if !In_array("2",[]string{"2","3"}) {
		t.Fatalf("util In_array has error")
	}
}

func TestHmacSha1(t *testing.T) {
	if HmacSha1("admin","testsalt") != "2e224647eeca047c0353eb9745c2f072dc5b2a17" {
		t.Fatalf("util HmacSha1 has error")
	}
}

func TestJwt(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.Init()
	JwtInit()


	token, err := GenerateToken(bson.NewObjectId(), "testname","testusername","testhash","testauth",60)

	if err != nil {
		t.Fatalf("gen jwt token has error:%v", err)
	}

	user, err := ParseToken(token)

	if err != nil {
		t.Fatalf("par jwt token has error:%v", err)
	}

	if user.Name != "testname" {
		t.Fatalf("jwt token data has error")
	}

	_, err = ParseToken("1d56a1")
	if err == nil {
		t.Fatalf("error token passed")
	}
}