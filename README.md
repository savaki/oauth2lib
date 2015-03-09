# oauth2lib

Pre-baked library for using golang.org/x/oauth2 with facebook, github, google, etc.  Compatible with any framework that uses the standard golang http.Handler

## Overview

1. You fill out an oauth2.Config sans Endpoint
2. Pass the config along with a callback function to the appropriate provider
3. You callback is invoked when the user authorizes a request
	- from your callback you can do whatever you want; redirect the user, create a new session for them, etc

## Simple Example

In this example, we redirect the user to /blah once they've successfully authorized via Facebook

```
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
```