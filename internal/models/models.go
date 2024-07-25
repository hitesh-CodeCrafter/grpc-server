package model

type User struct {
	ID       int32   `json:"id"`
	FullName string  `json:"full_name"`
	City     string  `json:"city"`
	Phone    int64   `json:"phone"`
	Height   float64 `json:"height"`
	Married  bool    `json:"Married"`
}

type SearchUsersRequest struct {
	Fname     string
	City      string
	Phone     int64
	MinHeight float64
	MaxHeight float64
	Married   *bool
}
