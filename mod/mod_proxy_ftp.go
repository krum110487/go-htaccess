// https://httpd.apache.org/docs/2.4/mod/mod_proxy_ftp.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_ftp struct{}

func (c *Mod_proxy_ftp) dirProxyFtpDirCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyFtpDirCharset directive
	return errors.New("ProxyFtpDirCharset is not yet implemented")
}

func (c *Mod_proxy_ftp) dirProxyFtpEscapeWildcards(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyFtpEscapeWildcards directive
	return errors.New("ProxyFtpEscapeWildcards is not yet implemented")
}

func (c *Mod_proxy_ftp) dirProxyFtpListOnWildcard(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyFtpListOnWildcard directive
	return errors.New("ProxyFtpListOnWildcard is not yet implemented")
}
