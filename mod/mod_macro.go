// https://httpd.apache.org/docs/2.4/mod/mod_macro.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_macro struct{}

func (Mod_macro) DirMacro(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: <Macro> directive
	return errors.New("<Macro> is not yet implemented")
}

func (Mod_macro) DirUndefMacro(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: UndefMacro directive
	return errors.New("UndefMacro is not yet implemented")
}

func (Mod_macro) DirUse(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Use directive
	return errors.New("Use is not yet implemented")
}
