// https://httpd.apache.org/docs/2.4/mod/mod_reflector.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_reflector struct{}

func (Mod_reflector) DirReflectorHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ReflectorHeader directive
	return errors.New("ReflectorHeader is not yet implemented")
}
