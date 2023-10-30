// https://httpd.apache.org/docs/2.4/mod/mod_imagemap.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_imagemap struct{}

func (c *Mod_imagemap) dirImapBase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ImapBase directive
	return errors.New("ImapBase is not yet implemented")
}

func (c *Mod_imagemap) dirImapDefault(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ImapDefault directive
	return errors.New("ImapDefault is not yet implemented")
}

func (c *Mod_imagemap) dirImapMenu(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ImapMenu directive
	return errors.New("ImapMenu is not yet implemented")
}
