// https://httpd.apache.org/docs/2.4/mod/mod_dialup.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dialup struct{}

func (Mod_dialup) DirModemStandard(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ModemStandard directive
	return errors.New("ModemStandard is not yet implemented")
}
