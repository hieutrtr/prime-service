package model

type Prime struct {
	HighestPrime uint32  `json:"highest_prime"`
}

type Input struct {
	Number uint32 `json:"number" validate:"required,numeric"`
}
