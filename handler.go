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
			Error: ErrNotAuthorized,
			w:     w,
		}
		h.callback(ctx)
		return
	}

	// otherwise, let's see what goodies we received
	t, err := h.exchange(oauth2.NoContext, code)
	ctx := &Context{
		Token:   t,
		Error:   err,
		Request: req,
		w:       w,
	}
	h.callback(ctx)
}

type CallbackFunc func(*Context)

type Context struct {
	Token   *oauth2.Token
	Error   error
	Request *http.Request
	w       http.ResponseWriter
}

func (c *Context) Header() http.Header {
	return c.w.Header()
}

func (c *Context) Write(b []byte) (int, error) {
	return c.w.Write(b)
}

func (c *Context) WriteHeader(code int) {
	c.w.WriteHeader(code)
}
