// https://httpd.apache.org/docs/2.4/mod/mod_rewrite.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_rewrite struct{}

func (c *Mod_rewrite) DirRewriteBase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RewriteBase directive
	return errors.New("RewriteBase is not yet implemented")
}

func (c *Mod_rewrite) DirRewriteCond(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RewriteCond directive
	return errors.New("RewriteCond is not yet implemented")
}

func (c *Mod_rewrite) DirRewriteEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RewriteEngine directive
	return errors.New("RewriteEngine is not yet implemented")
}

func (c *Mod_rewrite) DirRewriteMap(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RewriteMap directive
	return errors.New("RewriteMap is not yet implemented")
}

func (c *Mod_rewrite) DirRewriteOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RewriteOptions directive
	return errors.New("RewriteOptions is not yet implemented")
}

func (c *Mod_rewrite) DirRewriteRule(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RewriteRule directive
	return errors.New("RewriteRule is not yet implemented")
}
