// https://httpd.apache.org/docs/2.4/mod/mod_info.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_info struct{}

func (c *Mod_info) dirAddModuleInfo(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddModuleInfo directive
	return errors.New("AddModuleInfo is not yet implemented")
}
