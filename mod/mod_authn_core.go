// https://httpd.apache.org/docs/2.4/mod/mod_authn_core.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authn_core struct{}

func (c *Mod_authn_core) DirAuthName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthName directive
	return errors.New("AuthName is not yet implemented")
}

func (c *Mod_authn_core) DirAuthnProviderAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <AuthnProviderAlias> directive
	return errors.New("<AuthnProviderAlias> is not yet implemented")
}

func (c *Mod_authn_core) DirAuthType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthType directive
	return errors.New("AuthType is not yet implemented")
}
