// https://httpd.apache.org/docs/2.4/mod/mod_sed.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_sed struct{}

func (c *Mod_sed) DirInputSed(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: InputSed directive
	return errors.New("InputSed is not yet implemented")
}

func (c *Mod_sed) DirOutputSed(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: OutputSed directive
	return errors.New("OutputSed is not yet implemented")
}
