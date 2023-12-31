// https://httpd.apache.org/docs/2.4/mod/mod_dav.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dav struct{}

func (Mod_dav) DirDav(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Dav directive
	return errors.New("Dav is not yet implemented")
}

func (Mod_dav) DirDavBasePath(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DavBasePath directive
	return errors.New("DavBasePath is not yet implemented")
}

func (Mod_dav) DirDavDepthInfinity(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DavDepthInfinity directive
	return errors.New("DavDepthInfinity is not yet implemented")
}

func (Mod_dav) DirDavMinTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DavMinTimeout directive
	return errors.New("DavMinTimeout is not yet implemented")
}
