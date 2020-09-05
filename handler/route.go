package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hieutrtr/prime-service/router/middleware"
)

func (h *Handler) RegisterV1(v *echo.Group, JWTSecret string) {
	jwtMiddleware := middleware.JWT([]byte(JWTSecret))
	prime := v.Group("/prime", jwtMiddleware)
	prime.POST("", h.HighestPrime)
}
