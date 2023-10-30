// https://httpd.apache.org/docs/2.4/mod/mod_request.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_request struct{}

func (c *Mod_request) dirKeptBodySize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: KeptBodySize directive
	return errors.New("KeptBodySize is not yet implemented")
}
