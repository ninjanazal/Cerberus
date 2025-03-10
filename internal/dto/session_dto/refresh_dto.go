package session_dto

// RefreshRequest represents the request payload for refreshing an access token.
// It contains the refresh token used to generate a new access token.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshData holds the data generated during the token refresh process.
// It includes the new access token and refresh token.
type RefreshData struct {
	AccessToken  string
	RefreshToken string
}

// RefreshResponse represents the response payload returned after a successful token refresh.
// It includes a success message, the new access token, and the new refresh token.
type RefreshResponse struct {
	Message      string `json:"message"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
