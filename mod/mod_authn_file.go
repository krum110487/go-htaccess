// https://httpd.apache.org/docs/2.4/mod/mod_authn_file.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authn_file struct{}

func (Mod_authn_file) DirAuthUserFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthUserFile directive
	return errors.New("AuthUserFile is not yet implemented")
}
