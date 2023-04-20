package response

type JwtResponse struct {
	Token     string `json:"token,omitempty"`
	ExpiresAt int64  `json:"expiresAt,omitempty"`
}
