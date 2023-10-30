// https://httpd.apache.org/docs/2.4/mod/mod_headers.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_headers struct{}

func (c *Mod_headers) dirHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Header directive
	return errors.New("Header is not yet implemented")
}

func (c *Mod_headers) dirRequestHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RequestHeader directive
	return errors.New("RequestHeader is not yet implemented")
}
