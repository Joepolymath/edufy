package adapters

import (
	fb "github.com/huandu/facebook/v2"
)

// FacebookResponseInfo represents the user information structure
type FacebookResponseInfo struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewFacebookAdapter() *Adapter {
	return &Adapter{}
}

// FacebookGetUserInfo retrieves user information using an access token
func (a *Adapter) FacebookGetUserInfo(accessToken string) (*FacebookResponseInfo, error) {

	res, err := fb.Get("/me", fb.Params{
		"fields":       cfg.FaceBookFields,
		"access_token": accessToken,
	})
	if err != nil {
		return nil, err
	}

	// Check for any errors in the response.
	if err := res.Err(); err != nil {
		return nil, err
	}
	res.GetField()

	var user FacebookResponseInfo
	if err := res.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
