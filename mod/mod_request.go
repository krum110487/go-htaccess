// https://httpd.apache.org/docs/2.4/mod/mod_request.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_request struct{}

func (Mod_request) DirKeptBodySize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: KeptBodySize directive
	return errors.New("KeptBodySize is not yet implemented")
}
