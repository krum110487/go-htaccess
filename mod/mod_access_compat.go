// https://httpd.apache.org/docs/2.4/mod/mod_access_compat.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_access_compat struct{}

func (Mod_access_compat) DirAllow(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Allow directive
	return errors.New("Allow is not yet implemented")
}

func (Mod_access_compat) DirDeny(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Deny directive
	return errors.New("Deny is not yet implemented")
}

func (Mod_access_compat) DirOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Order directive
	return errors.New("Order is not yet implemented")
}

func (Mod_access_compat) DirSatisfy(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Satisfy directive
	return errors.New("Satisfy is not yet implemented")
}
