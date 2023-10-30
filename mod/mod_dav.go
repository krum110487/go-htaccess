// https://httpd.apache.org/docs/2.4/mod/mod_dav.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dav struct{}

func (c *Mod_dav) dirDav(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Dav directive
	return errors.New("Dav is not yet implemented")
}

func (c *Mod_dav) dirDavBasePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DavBasePath directive
	return errors.New("DavBasePath is not yet implemented")
}

func (c *Mod_dav) dirDavDepthInfinity(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DavDepthInfinity directive
	return errors.New("DavDepthInfinity is not yet implemented")
}

func (c *Mod_dav) dirDavMinTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DavMinTimeout directive
	return errors.New("DavMinTimeout is not yet implemented")
}
