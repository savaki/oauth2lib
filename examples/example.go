package main

import (
	"fmt"
	"net/http"

	"github.com/savaki/oauth2lib"
	"golang.org/x/oauth2"
)

func main() {
	config := &oauth2.Config{
		ClientID:     "your-facebook-app-id",
		ClientSecret: "your-facebook-app-secret",
		RedirectURL:  "your-oauth2-callback-endpoint",
	}

	handler := oauth2lib.Facebook(config, authorize)
	http.Handle("/oauth2/facebook", handler)

	http.ListenAndServe(":8000", nil)
}

func authorize(c *oauth2lib.Context) {
	fmt.Println("access token is", c.Token.AccessToken)
	http.Redirect(c.Response, c.Request, "/blah", http.StatusTemporaryRedirect)
}
