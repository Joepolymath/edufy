package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Adapter struct {
	httpClient *http.Client
}

func NewGoogleAdapter() *Adapter {
	// Initialize the http client here
	httpClient := &http.Client{
		// Configure HTTP client settings as needed
	}
	return &Adapter{
		httpClient: httpClient,
	}
}

// GoogleResponseInfo User represents the user information structure
type GoogleResponseInfo struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	FirstName       string `json:"given_name"`
	LastName        string `json:"family_name"`
	IsEmailVerified bool   `json:"verified_email"`
	ProfileImage    string `json:"picture"`
}

// GoogleGetUserInfo retrieves user information using an access token
func (a *Adapter) GoogleGetUserInfo(accessToken string) (*GoogleResponseInfo, error) {

	// OAuth2 configuration
	var googleOAuthConfig = oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.GoogleCallbackUrl,
		Scopes:       []string{cfg.GoogleScopesEmailUrl, cfg.GoogleScopesProfileUrl},
		Endpoint:     google.Endpoint,
	}
	client := googleOAuthConfig.Client(oauth2.NoContext, &oauth2.Token{AccessToken: accessToken})

	resp, err := client.Get(cfg.GoogleOauthUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user information: %s", resp.Status)
	}

	var user GoogleResponseInfo
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
