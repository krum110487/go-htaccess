// https://httpd.apache.org/docs/2.4/mod/mod_authz_groupfile.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authz_groupfile struct{}

func (c *Mod_authz_groupfile) DirAuthGroupFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthGroupFile directive
	return errors.New("AuthGroupFile is not yet implemented")
}
