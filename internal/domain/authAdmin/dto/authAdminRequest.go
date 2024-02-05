package dto

type AdminLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
