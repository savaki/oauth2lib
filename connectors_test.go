package oauth2lib

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func TestConnectors(t *testing.T) {
	fakeCode := "blah"
	fakeToken := &oauth2.Token{
		AccessToken: "blah blah",
		TokenType:   "blah blah blah",
	}

	Convey("Given an oauth2 connector", t, func() {
		var ctx *Context
		h := Google(&oauth2.Config{}, func(c *Context) {
			ctx = c
		})

		Convey("When a valid code is received via #ServeHTTP", func() {
			req, _ := http.NewRequest("GET", "http://acme.co?code="+fakeCode, nil)

			// fake oauth2.Config#Exchange
			h.(*handler).exchange = func(c context.Context, code string) (*oauth2.Token, error) {
				So(code, ShouldEqual, fakeCode)
				return fakeToken, nil
			}

			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)

			Convey("Then I expect the fakeToken to be received", func() {
				So(ctx.Token, ShouldResemble, fakeToken)
			})

			Convey("And I expect no errors to be returned", func() {
				So(ctx.Error, ShouldBeNil)
			})

			Convey("And I expect the response write to be set", func() {
				So(ctx.Response, ShouldEqual, w)
			})
		})

		Convey("When a no code is received via #ServeHTTP", func() {
			req, _ := http.NewRequest("GET", "http://acme.co", nil)

			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)

			Convey("Then I expect no fakeToken to be received", func() {
				So(ctx.Token, ShouldBeNil)
			})

			Convey("And I expect an ErrNotAuthorized to be returned", func() {
				So(ctx.Error, ShouldEqual, ErrNotAuthorized)
			})

			Convey("And I expect the response write to be set", func() {
				So(ctx.Response, ShouldEqual, w)
			})
		})
	})
}
