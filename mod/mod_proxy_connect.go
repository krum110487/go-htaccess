// https://httpd.apache.org/docs/2.4/mod/mod_proxy_connect.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_connect struct{}

func (c *Mod_proxy_connect) dirAllowCONNECT(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AllowCONNECT directive
	return errors.New("AllowCONNECT is not yet implemented")
}
