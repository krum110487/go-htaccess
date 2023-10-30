// https://httpd.apache.org/docs/2.4/mod/mod_setenvif.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_setenvif struct{}

func (c *Mod_setenvif) dirBrowserMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BrowserMatch directive
	return errors.New("BrowserMatch is not yet implemented")
}

func (c *Mod_setenvif) dirBrowserMatchNoCase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BrowserMatchNoCase directive
	return errors.New("BrowserMatchNoCase is not yet implemented")
}

func (c *Mod_setenvif) dirSetEnvIf(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetEnvIf directive
	return errors.New("SetEnvIf is not yet implemented")
}

func (c *Mod_setenvif) dirSetEnvIfExpr(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetEnvIfExpr directive
	return errors.New("SetEnvIfExpr is not yet implemented")
}

func (c *Mod_setenvif) dirSetEnvIfNoCase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetEnvIfNoCase directive
	return errors.New("SetEnvIfNoCase is not yet implemented")
}
