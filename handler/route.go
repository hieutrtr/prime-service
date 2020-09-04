package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterV1(v *echo.Group) {
	prime := v.Group("/prime")
	prime.POST("", h.HighestPrime)
}
