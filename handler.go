package oauth2lib

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	ErrNotAuthorized = fmt.Errorf("user didn't authorize request")
)

type handler struct {
	callback CallbackFunc
	config   *oauth2.Config
	exchange func(context.Context, string) (*oauth2.Token, error)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")

	// no code?  user must not have authorized the request
	if code == "" {
		ctx := &Context{
			Error:    ErrNotAuthorized,
			Response: w,
		}
		h.callback(ctx)
		return
	}

	// otherwise, let's see what goodies we received
	t, err := h.exchange(oauth2.NoContext, code)
	ctx := &Context{
		Token:    t,
		Error:    err,
		Request:  req,
		Response: w,
	}
	h.callback(ctx)
}

// CallbackFunc defines the user callback that gets invoked once the user
// authorizes (or doesn't) the application
type CallbackFunc func(*Context)

// Context holds the relevant values for the CallbackFunc
type Context struct {
	// the oauth2 authorization token
	Token *oauth2.Token

	// holds any error that occurrs while attempting to retrieve the oauth token
	// if the user does not authorize the request, that too is considered an
	// err and returns oauth2lib.ErrNotAuthorized
	Error error

	// the original request received by the user.  only provided to support methods
	// like http.Redirect which require the original request
	Request *http.Request

	// the original response the user may use for redirects and the like
	Response http.ResponseWriter
}
