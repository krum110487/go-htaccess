// https://httpd.apache.org/docs/2.4/mod/mod_userdir.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_userdir struct{}

func (Mod_userdir) DirUserDir(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: UserDir directive
	return errors.New("UserDir is not yet implemented")
}
