// https://httpd.apache.org/docs/2.4/mod/mod_dumpio.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dumpio struct{}

func (Mod_dumpio) DirDumpIOInput(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DumpIOInput directive
	return errors.New("DumpIOInput is not yet implemented")
}

func (Mod_dumpio) DirDumpIOOutput(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DumpIOOutput directive
	return errors.New("DumpIOOutput is not yet implemented")
}
