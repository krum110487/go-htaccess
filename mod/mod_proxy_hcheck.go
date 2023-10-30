// https://httpd.apache.org/docs/2.4/mod/mod_proxy_hcheck.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_hcheck struct{}

func (c *Mod_proxy_hcheck) DirProxyHCExpr(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHCExpr directive
	return errors.New("ProxyHCExpr is not yet implemented")
}

func (c *Mod_proxy_hcheck) DirProxyHCTemplate(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHCTemplate directive
	return errors.New("ProxyHCTemplate is not yet implemented")
}

func (c *Mod_proxy_hcheck) DirProxyHCTPsize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHCTPsize directive
	return errors.New("ProxyHCTPsize is not yet implemented")
}
