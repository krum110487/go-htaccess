// https://httpd.apache.org/docs/2.4/mod/mod_dav_fs.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dav_fs struct{}

func (c *Mod_dav_fs) dirDavLockDB(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DavLockDB directive
	return errors.New("DavLockDB is not yet implemented")
}

func (c *Mod_dav_fs) dirDavLockDiscovery(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DavLockDiscovery directive
	return errors.New("DavLockDiscovery is not yet implemented")
}
