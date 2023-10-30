// https://httpd.apache.org/docs/2.4/mod/mod_session_cookie.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_session_cookie struct{}

func (c *Mod_session_cookie) dirSessionCookieName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionCookieName directive
	return errors.New("SessionCookieName is not yet implemented")
}

func (c *Mod_session_cookie) dirSessionCookieName2(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionCookieName2 directive
	return errors.New("SessionCookieName2 is not yet implemented")
}

func (c *Mod_session_cookie) dirSessionCookieRemove(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionCookieRemove directive
	return errors.New("SessionCookieRemove is not yet implemented")
}
