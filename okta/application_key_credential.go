package okta

import "time"

type ApplicationKeyCredential struct {
	Created    *time.Time `json:"created,omitempty"`
	E          string     `json:"e,omitempty"`
	Expires_at *time.Time `json:"expiresAt,omitempty"`
	Kid        string     `json:"kid,omitempty"`
	Kty        string     `json:"kty,omitempty"`
	N          string     `json:"n,omitempty"`
	Use        string     `json:"use,omitempty"`
	X5c        []string   `json:"x5c,omitempty"`
	X5t_s256   string     `json:"x5t#S256,omitempty"`
}
