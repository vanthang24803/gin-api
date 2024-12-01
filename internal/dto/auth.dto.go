package dto

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Payload struct {
	Id       string   `json:"id"`
	Email    string   `json:"email"`
	FullName string   `json:"fullName"`
	Roles    []string `json:"roles"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
