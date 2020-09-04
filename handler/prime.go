package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) HighestPrime(c echo.Context) (err error) {
	return c.JSON(http.StatusOK,"something")
}
