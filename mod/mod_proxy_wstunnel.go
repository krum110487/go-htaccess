// https://httpd.apache.org/docs/2.4/mod/mod_proxy_wstunnel.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_wstunnel struct{}

func (c *Mod_proxy_wstunnel) dirProxyWebsocketFallbackToProxyHttp(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyWebsocketFallbackToProxyHttp directive
	return errors.New("ProxyWebsocketFallbackToProxyHttp is not yet implemented")
}
