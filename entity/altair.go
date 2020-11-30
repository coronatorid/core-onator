package entity

import "time"

// GrantTokenRequest used for altair grant token request
type GrantTokenRequest struct {
	ClientUID       string `json:"client_uid"`
	ClientSecret    string `json:"client_secret"`
	ResourceOwnerID int    `json:"resource_owner_id"`
	ResponseType    string `json:"response_type"`
	RedirectURI     string `json:"redirect_uri"`
	Scopes          string `json:"scopes"`
}

// OauthAccessToken represent Altair's response for oauth access token
type OauthAccessToken struct {
	Data struct {
		ID                 int       `json:"id"`
		OauthApplicationID int       `json:"oauth_application_id"`
		ResourceOwnerID    int       `json:"resource_owner_id"`
		Token              string    `json:"token"`
		Scopes             string    `json:"scopes"`
		ExpiresIn          int       `json:"expires_in"`
		RedirectURI        string    `json:"redirect_uri"`
		CreatedAt          time.Time `json:"created_at"`
	} `json:"data"`
}

// AltairError format error got from Altair
type AltairError struct {
	Error []struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"errors"`
}
