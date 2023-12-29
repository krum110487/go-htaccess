// https://httpd.apache.org/docs/2.4/mod/mod_so.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_so struct{}

func (Mod_so) DirLoadFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: LoadFile directive
	return errors.New("LoadFile is not yet implemented")
}

func (Mod_so) DirLoadModule(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: LoadModule directive
	return errors.New("LoadModule is not yet implemented")
}
