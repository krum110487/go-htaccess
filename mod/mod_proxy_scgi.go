// https://httpd.apache.org/docs/2.4/mod/mod_proxy_scgi.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_scgi struct{}

func (Mod_proxy_scgi) DirProxySCGIInternalRedirect(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProxySCGIInternalRedirect directive
	return errors.New("ProxySCGIInternalRedirect is not yet implemented")
}

func (Mod_proxy_scgi) DirProxySCGISendfile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProxySCGISendfile directive
	return errors.New("ProxySCGISendfile is not yet implemented")
}
