// https://httpd.apache.org/docs/2.4/mod/mod_authn_socache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authn_socache struct{}

func (c *Mod_authn_socache) DirAuthnCacheContext(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthnCacheContext directive
	return errors.New("AuthnCacheContext is not yet implemented")
}

func (c *Mod_authn_socache) DirAuthnCacheEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthnCacheEnable directive
	return errors.New("AuthnCacheEnable is not yet implemented")
}

func (c *Mod_authn_socache) DirAuthnCacheProvideFor(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthnCacheProvideFor directive
	return errors.New("AuthnCacheProvideFor is not yet implemented")
}

func (c *Mod_authn_socache) DirAuthnCacheSOCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthnCacheSOCache directive
	return errors.New("AuthnCacheSOCache is not yet implemented")
}

func (c *Mod_authn_socache) DirAuthnCacheTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthnCacheTimeout directive
	return errors.New("AuthnCacheTimeout is not yet implemented")
}
