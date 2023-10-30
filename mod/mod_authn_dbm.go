// https://httpd.apache.org/docs/2.4/mod/mod_authn_dbm.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authn_dbm struct{}

func (c *Mod_authn_dbm) dirAuthDBMType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDBMType directive
	return errors.New("AuthDBMType is not yet implemented")
}

func (c *Mod_authn_dbm) dirAuthDBMUserFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDBMUserFile directive
	return errors.New("AuthDBMUserFile is not yet implemented")
}
