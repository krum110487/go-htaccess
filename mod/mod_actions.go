// https://httpd.apache.org/docs/2.4/mod/mod_actions.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_actions struct{}

func (c *Mod_actions) DirAction(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Action directive
	return errors.New("Action is not yet implemented")
}

func (c *Mod_actions) DirScript(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Script directive
	return errors.New("Script is not yet implemented")
}
