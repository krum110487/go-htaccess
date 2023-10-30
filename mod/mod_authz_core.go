// https://httpd.apache.org/docs/2.4/mod/mod_authz_core.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authz_core struct{}

func (c *Mod_authz_core) dirAuthMerging(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthMerging directive
	return errors.New("AuthMerging is not yet implemented")
}

func (c *Mod_authz_core) dirAuthzProviderAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <AuthzProviderAlias> directive
	return errors.New("<AuthzProviderAlias> is not yet implemented")
}

func (c *Mod_authz_core) dirAuthzSendForbiddenOnFailure(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthzSendForbiddenOnFailure directive
	return errors.New("AuthzSendForbiddenOnFailure is not yet implemented")
}

func (c *Mod_authz_core) dirRequire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Require directive
	return errors.New("Require is not yet implemented")
}

func (c *Mod_authz_core) dirRequireAll(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <RequireAll> directive
	return errors.New("<RequireAll> is not yet implemented")
}

func (c *Mod_authz_core) dirRequireAny(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <RequireAny> directive
	return errors.New("<RequireAny> is not yet implemented")
}

func (c *Mod_authz_core) dirRequireNone(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <RequireNone> directive
	return errors.New("<RequireNone> is not yet implemented")
}
