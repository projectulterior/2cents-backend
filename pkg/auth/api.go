package auth

type TokenResponse struct {
	Auth    string `json:"auth_token"`
	Refresh string `json:"refresh_token"`
}
