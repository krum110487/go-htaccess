// https://httpd.apache.org/docs/2.4/mod/mod_proxy_fcgi.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_fcgi struct{}

func (Mod_proxy_fcgi) DirProxyFCGIBackendType(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProxyFCGIBackendType directive
	return errors.New("ProxyFCGIBackendType is not yet implemented")
}

func (Mod_proxy_fcgi) DirProxyFCGISetEnvIf(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProxyFCGISetEnvIf directive
	return errors.New("ProxyFCGISetEnvIf is not yet implemented")
}
