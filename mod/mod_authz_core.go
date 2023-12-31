// https://httpd.apache.org/docs/2.4/mod/mod_authz_core.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authz_core struct{}

func (Mod_authz_core) DirAuthMerging(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthMerging directive
	return errors.New("AuthMerging is not yet implemented")
}

func (Mod_authz_core) DirAuthzProviderAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: <AuthzProviderAlias> directive
	return errors.New("<AuthzProviderAlias> is not yet implemented")
}

func (Mod_authz_core) DirAuthzSendForbiddenOnFailure(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthzSendForbiddenOnFailure directive
	return errors.New("AuthzSendForbiddenOnFailure is not yet implemented")
}

func (Mod_authz_core) DirRequire(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Require directive
	return errors.New("Require is not yet implemented")
}

func (Mod_authz_core) DirRequireAll(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: <RequireAll> directive
	return errors.New("<RequireAll> is not yet implemented")
}

func (Mod_authz_core) DirRequireAny(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: <RequireAny> directive
	return errors.New("<RequireAny> is not yet implemented")
}

func (Mod_authz_core) DirRequireNone(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: <RequireNone> directive
	return errors.New("<RequireNone> is not yet implemented")
}
