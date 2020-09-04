package handler

import "prime-service/store"

type Handler struct {
	primeCache *store.PrimeCache
}

func NewHandler(primeCache *store.PrimeCache) *Handler {
	return &Handler{
		primeCache,
	}
}
