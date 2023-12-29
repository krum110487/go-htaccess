// https://httpd.apache.org/docs/2.4/mod/mod_buffer.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_buffer struct{}

func (Mod_buffer) DirBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: BufferSize directive
	return errors.New("BufferSize is not yet implemented")
}
