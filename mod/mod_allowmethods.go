// https://httpd.apache.org/docs/2.4/mod/mod_allowmethods.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_allowmethods struct{}

func (Mod_allowmethods) DirAllowMethods(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AllowMethods directive
	return errors.New("AllowMethods is not yet implemented")
}
