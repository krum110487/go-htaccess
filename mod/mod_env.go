// https://httpd.apache.org/docs/2.4/mod/mod_env.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_env struct{}

func (c *Mod_env) dirPassEnv(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: PassEnv directive
	return errors.New("PassEnv is not yet implemented")
}

func (c *Mod_env) dirSetEnv(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetEnv directive
	return errors.New("SetEnv is not yet implemented")
}

func (c *Mod_env) dirUnsetEnv(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UnsetEnv directive
	return errors.New("UnsetEnv is not yet implemented")
}