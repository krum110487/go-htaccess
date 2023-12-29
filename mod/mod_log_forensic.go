// https://httpd.apache.org/docs/2.4/mod/mod_log_forensic.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_log_forensic struct{}

func (Mod_log_forensic) DirForensicLog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ForensicLog directive
	return errors.New("ForensicLog is not yet implemented")
}
