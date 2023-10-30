// https://httpd.apache.org/docs/2.4/mod/mod_charset_lite.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_charset_lite struct{}

func (c *Mod_charset_lite) DirCharsetDefault(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CharsetDefault directive
	return errors.New("CharsetDefault is not yet implemented")
}

func (c *Mod_charset_lite) DirCharsetOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CharsetOptions directive
	return errors.New("CharsetOptions is not yet implemented")
}

func (c *Mod_charset_lite) DirCharsetSourceEnc(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CharsetSourceEnc directive
	return errors.New("CharsetSourceEnc is not yet implemented")
}
