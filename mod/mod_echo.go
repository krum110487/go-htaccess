// https://httpd.apache.org/docs/2.4/mod/mod_echo.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_echo struct{}

func (Mod_echo) DirProtocolEcho(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProtocolEcho directive
	return errors.New("ProtocolEcho is not yet implemented")
}
