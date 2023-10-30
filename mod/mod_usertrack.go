// https://httpd.apache.org/docs/2.4/mod/mod_usertrack.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_usertrack struct{}

func (c *Mod_usertrack) DirCookieDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieDomain directive
	return errors.New("CookieDomain is not yet implemented")
}

func (c *Mod_usertrack) DirCookieExpires(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieExpires directive
	return errors.New("CookieExpires is not yet implemented")
}

func (c *Mod_usertrack) DirCookieHTTPOnly(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieHTTPOnly directive
	return errors.New("CookieHTTPOnly is not yet implemented")
}

func (c *Mod_usertrack) DirCookieName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieName directive
	return errors.New("CookieName is not yet implemented")
}

func (c *Mod_usertrack) DirCookieSameSite(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieSameSite directive
	return errors.New("CookieSameSite is not yet implemented")
}

func (c *Mod_usertrack) DirCookieSecure(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieSecure directive
	return errors.New("CookieSecure is not yet implemented")
}

func (c *Mod_usertrack) DirCookieStyle(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieStyle directive
	return errors.New("CookieStyle is not yet implemented")
}

func (c *Mod_usertrack) DirCookieTracking(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CookieTracking directive
	return errors.New("CookieTracking is not yet implemented")
}
