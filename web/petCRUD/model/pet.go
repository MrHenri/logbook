package model

type Pet struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Species string `json:"species"`
	Sex     string `json:"sex"`
	Birth   string `json:"birth"`
	Death   string `json:"death"`
}
