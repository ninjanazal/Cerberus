package postgres_dto

// RegisterResponse represents the response structure for a successful user registration.
//
// Fields:
//   - Message: A string containing a success message or additional information about the registration.
//   - UserId: An unsigned integer representing the unique identifier of the newly registered user.
type RegisterResponse struct {
	Message string `json:"message"`
	UserId  uint   `json:"user_id"`
}

// RegisterRequest represents the request structure for user registration.
//
// Fields:
//   - Email: A string containing the email address of the user to be registered.
//   - Name: A string containing the name of the user to be registered.
//   - Password: A string containing the password for the new user account.
type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
