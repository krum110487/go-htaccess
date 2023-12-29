// https://httpd.apache.org/docs/2.4/mod/mod_authz_dbd.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authz_dbd struct{}

func (Mod_authz_dbd) DirAuthzDBDLoginToReferer(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthzDBDLoginToReferer directive
	return errors.New("AuthzDBDLoginToReferer is not yet implemented")
}

func (Mod_authz_dbd) DirAuthzDBDQuery(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthzDBDQuery directive
	return errors.New("AuthzDBDQuery is not yet implemented")
}

func (Mod_authz_dbd) DirAuthzDBDRedirectQuery(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthzDBDRedirectQuery directive
	return errors.New("AuthzDBDRedirectQuery is not yet implemented")
}
