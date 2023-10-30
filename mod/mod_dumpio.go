// https://httpd.apache.org/docs/2.4/mod/mod_dumpio.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dumpio struct{}

func (c *Mod_dumpio) dirDumpIOInput(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DumpIOInput directive
	return errors.New("DumpIOInput is not yet implemented")
}

func (c *Mod_dumpio) dirDumpIOOutput(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DumpIOOutput directive
	return errors.New("DumpIOOutput is not yet implemented")
}
