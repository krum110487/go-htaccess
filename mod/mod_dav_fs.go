// https://httpd.apache.org/docs/2.4/mod/mod_dav_fs.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dav_fs struct{}

func (Mod_dav_fs) DirDavLockDB(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DavLockDB directive
	return errors.New("DavLockDB is not yet implemented")
}

func (Mod_dav_fs) DirDavLockDiscovery(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DavLockDiscovery directive
	return errors.New("DavLockDiscovery is not yet implemented")
}
