// https://httpd.apache.org/docs/2.4/mod/mod_proxy_scgi.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_scgi struct{}

func (c *Mod_proxy_scgi) dirProxySCGIInternalRedirect(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxySCGIInternalRedirect directive
	return errors.New("ProxySCGIInternalRedirect is not yet implemented")
}

func (c *Mod_proxy_scgi) dirProxySCGISendfile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxySCGISendfile directive
	return errors.New("ProxySCGISendfile is not yet implemented")
}
