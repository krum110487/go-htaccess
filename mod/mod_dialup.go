// https://httpd.apache.org/docs/2.4/mod/mod_dialup.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dialup struct{}

func (c *Mod_dialup) dirModemStandard(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ModemStandard directive
	return errors.New("ModemStandard is not yet implemented")
}
