// https://httpd.apache.org/docs/2.4/mod/mod_suexec.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_suexec struct{}

func (Mod_suexec) DirSuexecUserGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SuexecUserGroup directive
	return errors.New("SuexecUserGroup is not yet implemented")
}
