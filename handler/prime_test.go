package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/hieutrtr/prime-service/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
)



func TestPrimeRequestSuccess(t *testing.T) {
	var reqJSON = `{"number": 567366}`
	req := httptest.NewRequest(echo.POST, "/api/v1/prime", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT(Issuer, JWTSecret))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.HighestPrime(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMapUint32(rec.Body.Bytes())
		assert.Equal(t, uint32(567323), m["highest_prime"])
	}
}

func TestPrimeRequestLessThanTwo(t *testing.T) {
	var reqJSON = `{"number": 1}`
	req := httptest.NewRequest(echo.POST, "/api/v1/prime", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT(Issuer, JWTSecret))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.HighestPrime(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusBadRequest, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "There's no prime less than 2", m["message"])
	}
}

func TestPrimeRequestInvalidInput(t *testing.T) {
	var reqJSON = `{"foo": 10}`
	req := httptest.NewRequest(echo.POST, "/api/v1/prime", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT(Issuer, JWTSecret))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.HighestPrime(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusBadRequest, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "Key: 'Input.Number' Error:Field validation for 'Number' failed on the 'required' tag", m["message"])
	}
}

func TestPrimeRequestOverLimitInput(t *testing.T) {
	var reqJSON = `{"number": 1000001}`
	req := httptest.NewRequest(echo.POST, "/api/v1/prime", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT(Issuer, JWTSecret))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.HighestPrime(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusBadRequest, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "The input is over limitation", m["message"])
	}
}

func TestUnauthorizedRequest(t *testing.T) {
	var reqJSON = `{"number": 1}`
	req := httptest.NewRequest(echo.POST, "/api/v1/prime", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	token := utils.GenerateJWT("Issuer", JWTSecret)
	fmt.Println(token)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT("Issuer", JWTSecret))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.HighestPrime(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusBadRequest, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		fmt.Println(m)
		assert.Equal(t, "User is unauthorized", m["message"])
	}
}
