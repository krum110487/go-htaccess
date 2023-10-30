// https://httpd.apache.org/docs/2.4/mod/mod_ident.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_ident struct{}

func (c *Mod_ident) DirIdentityCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IdentityCheck directive
	return errors.New("IdentityCheck is not yet implemented")
}

func (c *Mod_ident) DirIdentityCheckTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IdentityCheckTimeout directive
	return errors.New("IdentityCheckTimeout is not yet implemented")
}
