package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
func fuckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := &util.Claims{
			Id:       "123",
			Name:     "admin",
			Username: "admin",
			Auth:     "admin",
			Hash:     "testhash",
		}

		session := sessions.Default(c)
		session.Set("NowUser", claims)

		c.Next()
	}
}

func testGinWithLogin() (server *gin.Engine) {
	gin.SetMode(gin.TestMode)

	config.Init()
	util.JwtInit()

	db.Connect()

	server = gin.New()
	server.Use(middlewares.Connect)

	// Support session
	store := cookie.NewStore([]byte(config.Get("secret").(string)))
	server.Use(sessions.Sessions("SESSION", store))
	server.Use(fuckJWT())

	return server
}

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
