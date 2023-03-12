package model

type Address struct {
	PinCode string `json:"pin_code"`
	City    string `json:"city"`
	State   string `json:"state"`
}
