// https://httpd.apache.org/docs/2.4/mod/mod_speling.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_speling struct{}

func (c *Mod_speling) dirCheckBasenameMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CheckBasenameMatch directive
	return errors.New("CheckBasenameMatch is not yet implemented")
}

func (c *Mod_speling) dirCheckCaseOnly(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CheckCaseOnly directive
	return errors.New("CheckCaseOnly is not yet implemented")
}

func (c *Mod_speling) dirCheckSpelling(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CheckSpelling directive
	return errors.New("CheckSpelling is not yet implemented")
}
