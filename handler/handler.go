package handler

import "github.com/hieutrtr/prime-service/store"

type Handler struct {
	primeCache *store.PrimeCache
}

func NewHandler(primeCache *store.PrimeCache) *Handler {
	return &Handler{
		primeCache,
	}
}
