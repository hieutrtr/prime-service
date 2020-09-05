package handler

import (
	"encoding/json"
	"github.com/hieutrtr/prime-service/router"
	"github.com/hieutrtr/prime-service/router/middleware"
	"github.com/hieutrtr/prime-service/store"

	"os"
	"testing"

	"github.com/labstack/echo/v4"
)
var (
	primeCache *store.PrimeCache
	h  *Handler
	e  *echo.Echo
	jwtMiddleware echo.MiddlewareFunc
	JWTSecret = "fortest"
	Issuer = "stably"
)
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	jwtMiddleware = middleware.JWT([]byte(JWTSecret))
	primeCache = store.NewPrimeCache(1000000)
	h = NewHandler(primeCache)
	e = router.New()
}


func responseMapUint32(b []byte) map[string]uint32 {
	var m map[string]uint32
	json.Unmarshal(b, &m)
	return m
}

func responseMap(b []byte) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m
}
