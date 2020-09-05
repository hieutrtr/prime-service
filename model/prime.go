package model

type Prime struct {
	HighestPrime uint32  `json:"highest_prime"`
}

type Input struct {
	Number uint32 `json:"number" validate:"required" example:"500000"`
}

type Error struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the error returned by an external dependency
}
