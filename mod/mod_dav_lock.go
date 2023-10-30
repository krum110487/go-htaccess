// https://httpd.apache.org/docs/2.4/mod/mod_dav_lock.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dav_lock struct{}

func (c *Mod_dav_lock) dirDavGenericLockDB(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DavGenericLockDB directive
	return errors.New("DavGenericLockDB is not yet implemented")
}
