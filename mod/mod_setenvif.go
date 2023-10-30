// https://httpd.apache.org/docs/2.4/mod/mod_setenvif.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_setenvif struct{}

func (c *Mod_setenvif) DirBrowserMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BrowserMatch directive
	return errors.New("BrowserMatch is not yet implemented")
}

func (c *Mod_setenvif) DirBrowserMatchNoCase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BrowserMatchNoCase directive
	return errors.New("BrowserMatchNoCase is not yet implemented")
}

func (c *Mod_setenvif) DirSetEnvIf(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetEnvIf directive
	return errors.New("SetEnvIf is not yet implemented")
}

func (c *Mod_setenvif) DirSetEnvIfExpr(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetEnvIfExpr directive
	return errors.New("SetEnvIfExpr is not yet implemented")
}

func (c *Mod_setenvif) DirSetEnvIfNoCase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetEnvIfNoCase directive
	return errors.New("SetEnvIfNoCase is not yet implemented")
}
