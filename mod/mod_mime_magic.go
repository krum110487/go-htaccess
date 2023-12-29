// https://httpd.apache.org/docs/2.4/mod/mod_mime_magic.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_mime_magic struct{}

func (Mod_mime_magic) DirMimeMagicFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MimeMagicFile directive
	return errors.New("MimeMagicFile is not yet implemented")
}
