// https://httpd.apache.org/docs/2.4/mod/mod_macro.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_macro struct{}

func (c *Mod_macro) dirMacro(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Macro> directive
	return errors.New("<Macro> is not yet implemented")
}

func (c *Mod_macro) dirUndefMacro(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UndefMacro directive
	return errors.New("UndefMacro is not yet implemented")
}

func (c *Mod_macro) dirUse(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Use directive
	return errors.New("Use is not yet implemented")
}
