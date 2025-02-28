package session_dto

// LoginRequest represents the structure of a login request.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the structure of a response to a login request.
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"toke"`
}
