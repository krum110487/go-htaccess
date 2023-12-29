// https://httpd.apache.org/docs/2.4/mod/mod_proxy_ftp.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_ftp struct{}

func (Mod_proxy_ftp) DirProxyFtpDirCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProxyFtpDirCharset directive
	return errors.New("ProxyFtpDirCharset is not yet implemented")
}

func (Mod_proxy_ftp) DirProxyFtpEscapeWildcards(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProxyFtpEscapeWildcards directive
	return errors.New("ProxyFtpEscapeWildcards is not yet implemented")
}

func (Mod_proxy_ftp) DirProxyFtpListOnWildcard(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ProxyFtpListOnWildcard directive
	return errors.New("ProxyFtpListOnWildcard is not yet implemented")
}
