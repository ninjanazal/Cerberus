package session_dto

// LoginRequest represents the data structure for a login request.
//
// Fields:
//   - Email: A string containing the user's email address. It is mapped to the "email" JSON field.
//   - Password: A string containing the user's password. It is mapped to the "password" JSON field.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginData represents the authentication tokens returned after a successful login.
//
// Fields:
//   - AccessToken: A string containing the JWT access token used for authenticating API requests.
//   - RefreshToken: A string containing the JWT refresh token used to obtain a new access token when it expires.
type LoginData struct {
	AccessToken  string
	RefreshToken string
}

// LoginResponse represents the structure of a response to a login request.
type LoginResponse struct {
	Message      string `json:"message"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
