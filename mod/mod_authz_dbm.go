// https://httpd.apache.org/docs/2.4/mod/mod_authz_dbm.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authz_dbm struct{}

func (c *Mod_authz_dbm) dirAuthDBMGroupFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDBMGroupFile directive
	return errors.New("AuthDBMGroupFile is not yet implemented")
}

func (c *Mod_authz_dbm) dirAuthzDBMType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthzDBMType directive
	return errors.New("AuthzDBMType is not yet implemented")
}
