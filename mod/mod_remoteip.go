// https://httpd.apache.org/docs/2.4/mod/mod_remoteip.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_remoteip struct{}

func (c *Mod_remoteip) DirRemoteIPHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPHeader directive
	return errors.New("RemoteIPHeader is not yet implemented")
}

func (c *Mod_remoteip) DirRemoteIPInternalProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPInternalProxy directive
	return errors.New("RemoteIPInternalProxy is not yet implemented")
}

func (c *Mod_remoteip) DirRemoteIPInternalProxyList(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPInternalProxyList directive
	return errors.New("RemoteIPInternalProxyList is not yet implemented")
}

func (c *Mod_remoteip) DirRemoteIPProxiesHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPProxiesHeader directive
	return errors.New("RemoteIPProxiesHeader is not yet implemented")
}

func (c *Mod_remoteip) DirRemoteIPProxyProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPProxyProtocol directive
	return errors.New("RemoteIPProxyProtocol is not yet implemented")
}

func (c *Mod_remoteip) DirRemoteIPProxyProtocolExceptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPProxyProtocolExceptions directive
	return errors.New("RemoteIPProxyProtocolExceptions is not yet implemented")
}

func (c *Mod_remoteip) DirRemoteIPTrustedProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPTrustedProxy directive
	return errors.New("RemoteIPTrustedProxy is not yet implemented")
}

func (c *Mod_remoteip) DirRemoteIPTrustedProxyList(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoteIPTrustedProxyList directive
	return errors.New("RemoteIPTrustedProxyList is not yet implemented")
}
