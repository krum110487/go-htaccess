// https://httpd.apache.org/docs/2.4/mod/mod_version.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_version struct{}

func (c *Mod_version) dirIfVersion(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfVersion> directive
	return errors.New("<IfVersion> is not yet implemented")
}
