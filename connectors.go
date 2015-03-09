package oauth2lib

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Special thanks to Weslley Andrade for the idea from the martini-contrib package

// Google returns a new Google OAuth 2.0 backend endpoint.
func Google(config *oauth2.Config, callback CallbackFunc) http.Handler {
	config.Endpoint = google.Endpoint
	return newHandler(config, callback)
}

// Github returns a new Github OAuth 2.0 backend endpoint.
func Github(config *oauth2.Config, callback CallbackFunc) http.Handler {
	config.Endpoint = oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	}
	return newHandler(config, callback)
}

// Facebook returns a new Facebook OAuth 2.0 backend endpoint.
func Facebook(config *oauth2.Config, callback CallbackFunc) http.Handler {
	config.Endpoint = oauth2.Endpoint{
		AuthURL:  "https://www.facebook.com/dialog/oauth",
		TokenURL: "https://graph.facebook.com/oauth/access_token",
	}
	return newHandler(config, callback)
}

// LinkedIn returns a new LinkedIn OAuth 2.0 backend endpoint.
func LinkedIn(config *oauth2.Config, callback CallbackFunc) http.Handler {
	config.Endpoint = oauth2.Endpoint{
		AuthURL:  "https://www.linkedin.com/uas/oauth2/authorization",
		TokenURL: "https://www.linkedin.com/uas/oauth2/accessToken",
	}
	return newHandler(config, callback)
}

func newHandler(config *oauth2.Config, callback CallbackFunc) http.Handler {
	return &handler{
		callback: callback,
		config:   config,
		exchange: config.Exchange,
	}
}
