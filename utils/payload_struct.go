package utils

// RegisterUserRequest represents the request payload for user registration
type RegisterUserRequest struct {
	FirstName       string `json:"firstName" binding:"required"`
	LastName        string `json:"lastName" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"ConfirmPassword" binding:"required"`
	CompanyName     string `json:"companyName" binding:"required"`
}

// AuthInput represents the request payload for login
type AuthInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// update the user firstName and lastName
type UpdateUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}


//Payment API payload
type PaymentInput struct {
    CardNumber string  `json:"card_number" binding:"required"`
    ExpiryDate string  `json:"expiry_date" binding:"required"`
    CVV        string  `json:"cvv" binding:"required"`
    Amount     float64 `json:"amount" binding:"required"`
}