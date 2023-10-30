// https://httpd.apache.org/docs/2.4/mod/mod_cgid.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cgid struct{}

func (c *Mod_cgid) dirCGIDScriptTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CGIDScriptTimeout directive
	return errors.New("CGIDScriptTimeout is not yet implemented")
}

func (c *Mod_cgid) dirScriptLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptLog directive
	return errors.New("ScriptLog is not yet implemented")
}

func (c *Mod_cgid) dirScriptLogBuffer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptLogBuffer directive
	return errors.New("ScriptLogBuffer is not yet implemented")
}

func (c *Mod_cgid) dirScriptLogLength(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptLogLength directive
	return errors.New("ScriptLogLength is not yet implemented")
}

func (c *Mod_cgid) dirScriptSock(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptSock directive
	return errors.New("ScriptSock is not yet implemented")
}
