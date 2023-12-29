// https://httpd.apache.org/docs/2.4/mod/mod_speling.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_speling struct{}

func (Mod_speling) DirCheckBasenameMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CheckBasenameMatch directive
	return errors.New("CheckBasenameMatch is not yet implemented")
}

func (Mod_speling) DirCheckCaseOnly(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CheckCaseOnly directive
	return errors.New("CheckCaseOnly is not yet implemented")
}

func (Mod_speling) DirCheckSpelling(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CheckSpelling directive
	return errors.New("CheckSpelling is not yet implemented")
}
