// https://httpd.apache.org/docs/2.4/mod/mod_cgid.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cgid struct{}

func (Mod_cgid) DirCGIDScriptTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CGIDScriptTimeout directive
	return errors.New("CGIDScriptTimeout is not yet implemented")
}

func (Mod_cgid) DirScriptLog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ScriptLog directive
	return errors.New("ScriptLog is not yet implemented")
}

func (Mod_cgid) DirScriptLogBuffer(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ScriptLogBuffer directive
	return errors.New("ScriptLogBuffer is not yet implemented")
}

func (Mod_cgid) DirScriptLogLength(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ScriptLogLength directive
	return errors.New("ScriptLogLength is not yet implemented")
}

func (Mod_cgid) DirScriptSock(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ScriptSock directive
	return errors.New("ScriptSock is not yet implemented")
}
