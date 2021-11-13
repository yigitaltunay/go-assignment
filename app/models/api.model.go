package models

import "time"

type InMemoryModel struct {
	Key        string    `json:"key,omitempty" binding:"required"`
	Value      Response  `json:"value,omitempty" binding:"required"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

type Config struct {
	Providers []string
	Port      string
}

type Response struct {
	Data           float64 `json:"data"`
	Currency       string  `json:"currency"`
	WhichProviders string  `json:"-"` // Don't send this to the client
}

type Currency struct {
	Eur float64 `json:"EUR"`
	Usd float64 `json:"USD"`
	Gbp float64 `json:"GBP"`
}
