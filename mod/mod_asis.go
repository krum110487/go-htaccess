// https://httpd.apache.org/docs/2.4/mod/mod_asis.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_asis struct{}

func (Mod_asis) DirUsage(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Usage directive
	return errors.New("Usage is not yet implemented")
}
