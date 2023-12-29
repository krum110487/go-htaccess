// https://httpd.apache.org/docs/2.4/mod/mod_log_debug.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_log_debug struct{}

func (Mod_log_debug) DirLogMessage(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: LogMessage directive
	return errors.New("LogMessage is not yet implemented")
}
