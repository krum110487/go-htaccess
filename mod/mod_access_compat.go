// https://httpd.apache.org/docs/2.4/mod/mod_access_compat.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_access_compat struct{}

func (c *Mod_access_compat) dirAllow(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Allow directive
	return errors.New("Allow is not yet implemented")
}

func (c *Mod_access_compat) dirDeny(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Deny directive
	return errors.New("Deny is not yet implemented")
}

func (c *Mod_access_compat) dirOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Order directive
	return errors.New("Order is not yet implemented")
}

func (c *Mod_access_compat) dirSatisfy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Satisfy directive
	return errors.New("Satisfy is not yet implemented")
}
