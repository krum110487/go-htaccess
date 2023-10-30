// https://httpd.apache.org/docs/2.4/mod/mod_expires.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_expires struct{}

func (c *Mod_expires) DirExpiresActive(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ExpiresActive directive
	return errors.New("ExpiresActive is not yet implemented")
}

func (c *Mod_expires) DirExpiresByType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ExpiresByType directive
	return errors.New("ExpiresByType is not yet implemented")
}

func (c *Mod_expires) DirExpiresDefault(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ExpiresDefault directive
	return errors.New("ExpiresDefault is not yet implemented")
}
