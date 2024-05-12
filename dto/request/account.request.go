package request

import "app/model"

type LoginGoogleRequest struct {
	Email         string     `json:"email"`
	EmailVerified bool       `json:"email_verified"`
	FamilyName    string     `json:"family_name"`
	GivenName     string     `json:"given_name"`
	Locale        string     `json:"locale"`
	Name          string     `json:"name"`
	Picture       string     `json:"picture"`
	Sub           string     `json:"sub"`
	Role          model.ROLE `json:"role"`
}
