// https://httpd.apache.org/docs/2.4/mod/mod_buffer.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_buffer struct{}

func (c *Mod_buffer) dirBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BufferSize directive
	return errors.New("BufferSize is not yet implemented")
}
