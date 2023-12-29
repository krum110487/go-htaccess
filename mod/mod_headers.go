// https://httpd.apache.org/docs/2.4/mod/mod_headers.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_headers struct{}

func (Mod_headers) DirHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Header directive
	return errors.New("Header is not yet implemented")
}

func (Mod_headers) DirRequestHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RequestHeader directive
	return errors.New("RequestHeader is not yet implemented")
}
