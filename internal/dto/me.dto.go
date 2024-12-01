package dto

type UpdateProfile struct {
	FirstName string `json:"firstName" binding:"required,min=2,max=50"`
	LastName  string `json:"lastName" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"required,email"`
}
