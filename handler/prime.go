package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"prime-service/model"
)

func traceBackHighestPrime(inputNumber uint32, marks []uint32) uint32 {
	halfInput := inputNumber/2-1
	prime := uint32(0)
	if marks[halfInput] == 0 {
		prime = halfInput*2+1
	} else {
		for i := halfInput-1; i > 0; i-- {
			if marks[i] == 0 {
				prime = i*2+1
				break
			}
		}
	}
	return prime
}

func (h *Handler) HighestPrime(c echo.Context) (err error) {
	input := new(model.Input)
	if err = c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	if err = c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	if input.Number > h.primeCache.Limit {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "The input is over limitation"))
	}
	prime := traceBackHighestPrime(input.Number, h.primeCache.Marks)
	result := model.Prime{
		HighestPrime: prime,
	}

	return c.JSON(http.StatusOK, result)
}
