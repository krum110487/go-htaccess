// https://httpd.apache.org/docs/2.4/mod/mod_proxy_express.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_express struct{}

func (c *Mod_proxy_express) DirProxyExpressDBMFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyExpressDBMFile directive
	return errors.New("ProxyExpressDBMFile is not yet implemented")
}

func (c *Mod_proxy_express) DirProxyExpressDBMType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyExpressDBMType directive
	return errors.New("ProxyExpressDBMType is not yet implemented")
}

func (c *Mod_proxy_express) DirProxyExpressEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyExpressEnable directive
	return errors.New("ProxyExpressEnable is not yet implemented")
}
