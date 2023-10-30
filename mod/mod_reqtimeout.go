// https://httpd.apache.org/docs/2.4/mod/mod_reqtimeout.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_reqtimeout struct{}

func (c *Mod_reqtimeout) DirRequestReadTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RequestReadTimeout directive
	return errors.New("RequestReadTimeout is not yet implemented")
}
