// https://httpd.apache.org/docs/2.4/mod/mod_authnz_fcgi.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authnz_fcgi struct{}

func (Mod_authnz_fcgi) DirAuthnzFcgiCheckAuthnProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthnzFcgiCheckAuthnProvider directive
	return errors.New("AuthnzFcgiCheckAuthnProvider is not yet implemented")
}

func (Mod_authnz_fcgi) DirAuthnzFcgiDefineProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthnzFcgiDefineProvider directive
	return errors.New("AuthnzFcgiDefineProvider is not yet implemented")
}
