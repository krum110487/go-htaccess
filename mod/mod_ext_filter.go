// https://httpd.apache.org/docs/2.4/mod/mod_ext_filter.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_ext_filter struct{}

func (c *Mod_ext_filter) dirExtFilterDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ExtFilterDefine directive
	return errors.New("ExtFilterDefine is not yet implemented")
}

func (c *Mod_ext_filter) dirExtFilterOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ExtFilterOptions directive
	return errors.New("ExtFilterOptions is not yet implemented")
}
