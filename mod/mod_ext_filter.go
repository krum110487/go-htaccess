// https://httpd.apache.org/docs/2.4/mod/mod_ext_filter.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_ext_filter struct{}

func (Mod_ext_filter) DirExtFilterDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ExtFilterDefine directive
	return errors.New("ExtFilterDefine is not yet implemented")
}

func (Mod_ext_filter) DirExtFilterOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ExtFilterOptions directive
	return errors.New("ExtFilterOptions is not yet implemented")
}
