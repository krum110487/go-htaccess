// https://httpd.apache.org/docs/2.4/mod/mod_cgi.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cgi struct{}

func (c *Mod_cgi) DirScriptLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptLog directive
	return errors.New("ScriptLog is not yet implemented")
}

func (c *Mod_cgi) DirScriptLogBuffer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptLogBuffer directive
	return errors.New("ScriptLogBuffer is not yet implemented")
}

func (c *Mod_cgi) DirScriptLogLength(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptLogLength directive
	return errors.New("ScriptLogLength is not yet implemented")
}
