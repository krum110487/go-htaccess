// https://httpd.apache.org/docs/2.4/mod/mod_nw_ssl.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_nw_ssl struct{}

func (Mod_nw_ssl) DirNWSSLTrustedCerts(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: NWSSLTrustedCerts directive
	return errors.New("NWSSLTrustedCerts is not yet implemented")
}

func (Mod_nw_ssl) DirNWSSLUpgradeable(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: NWSSLUpgradeable directive
	return errors.New("NWSSLUpgradeable is not yet implemented")
}

func (Mod_nw_ssl) DirSecureListen(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SecureListen directive
	return errors.New("SecureListen is not yet implemented")
}
