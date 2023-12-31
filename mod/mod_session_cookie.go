// https://httpd.apache.org/docs/2.4/mod/mod_session_cookie.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_session_cookie struct{}

func (Mod_session_cookie) DirSessionCookieName(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionCookieName directive
	return errors.New("SessionCookieName is not yet implemented")
}

func (Mod_session_cookie) DirSessionCookieName2(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionCookieName2 directive
	return errors.New("SessionCookieName2 is not yet implemented")
}

func (Mod_session_cookie) DirSessionCookieRemove(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionCookieRemove directive
	return errors.New("SessionCookieRemove is not yet implemented")
}
