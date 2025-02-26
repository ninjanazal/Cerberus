package auth_dto

// ChangePasswordRequest represents the data structure for a password change request.
type ChangePasswordRequest struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// ChangePasswordResponse represents the data structure for a password change response.
type ChangePasswordResponse struct {
	Message string `json:"message"`
}
