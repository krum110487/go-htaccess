// https://httpd.apache.org/docs/2.4/mod/mod_proxy.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy struct{}

func (c *Mod_proxy) dirBalancerGrowth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerGrowth directive
	return errors.New("BalancerGrowth is not yet implemented")
}

func (c *Mod_proxy) dirBalancerInherit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerInherit directive
	return errors.New("BalancerInherit is not yet implemented")
}

func (c *Mod_proxy) dirBalancerMember(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerMember directive
	return errors.New("BalancerMember is not yet implemented")
}

func (c *Mod_proxy) dirBalancerPersist(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerPersist directive
	return errors.New("BalancerPersist is not yet implemented")
}

func (c *Mod_proxy) dirNoProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: NoProxy directive
	return errors.New("NoProxy is not yet implemented")
}

func (c *Mod_proxy) dirProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Proxy> directive
	return errors.New("<Proxy> is not yet implemented")
}

func (c *Mod_proxy) dirProxy100Continue(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Proxy100Continue directive
	return errors.New("Proxy100Continue is not yet implemented")
}

func (c *Mod_proxy) dirProxyAddHeaders(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyAddHeaders directive
	return errors.New("ProxyAddHeaders is not yet implemented")
}

func (c *Mod_proxy) dirProxyBadHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyBadHeader directive
	return errors.New("ProxyBadHeader is not yet implemented")
}

func (c *Mod_proxy) dirProxyBlock(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyBlock directive
	return errors.New("ProxyBlock is not yet implemented")
}

func (c *Mod_proxy) dirProxyDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyDomain directive
	return errors.New("ProxyDomain is not yet implemented")
}

func (c *Mod_proxy) dirProxyErrorOverride(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyErrorOverride directive
	return errors.New("ProxyErrorOverride is not yet implemented")
}

func (c *Mod_proxy) dirProxyIOBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyIOBufferSize directive
	return errors.New("ProxyIOBufferSize is not yet implemented")
}

func (c *Mod_proxy) dirProxyMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <ProxyMatch> directive
	return errors.New("<ProxyMatch> is not yet implemented")
}

func (c *Mod_proxy) dirProxyMaxForwards(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyMaxForwards directive
	return errors.New("ProxyMaxForwards is not yet implemented")
}

func (c *Mod_proxy) dirProxyPass(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPass directive
	return errors.New("ProxyPass is not yet implemented")
}

func (c *Mod_proxy) dirProxyPassInherit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassInherit directive
	return errors.New("ProxyPassInherit is not yet implemented")
}

func (c *Mod_proxy) dirProxyPassInterpolateEnv(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassInterpolateEnv directive
	return errors.New("ProxyPassInterpolateEnv is not yet implemented")
}

func (c *Mod_proxy) dirProxyPassMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassMatch directive
	return errors.New("ProxyPassMatch is not yet implemented")
}

func (c *Mod_proxy) dirProxyPassReverse(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassReverse directive
	return errors.New("ProxyPassReverse is not yet implemented")
}

func (c *Mod_proxy) dirProxyPassReverseCookieDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassReverseCookieDomain directive
	return errors.New("ProxyPassReverseCookieDomain is not yet implemented")
}

func (c *Mod_proxy) dirProxyPassReverseCookiePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassReverseCookiePath directive
	return errors.New("ProxyPassReverseCookiePath is not yet implemented")
}

func (c *Mod_proxy) dirProxyPreserveHost(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPreserveHost directive
	return errors.New("ProxyPreserveHost is not yet implemented")
}

func (c *Mod_proxy) dirProxyReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyReceiveBufferSize directive
	return errors.New("ProxyReceiveBufferSize is not yet implemented")
}

func (c *Mod_proxy) dirProxyRemote(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyRemote directive
	return errors.New("ProxyRemote is not yet implemented")
}

func (c *Mod_proxy) dirProxyRemoteMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyRemoteMatch directive
	return errors.New("ProxyRemoteMatch is not yet implemented")
}

func (c *Mod_proxy) dirProxyRequests(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyRequests directive
	return errors.New("ProxyRequests is not yet implemented")
}

func (c *Mod_proxy) dirProxySet(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxySet directive
	return errors.New("ProxySet is not yet implemented")
}

func (c *Mod_proxy) dirProxySourceAddress(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxySourceAddress directive
	return errors.New("ProxySourceAddress is not yet implemented")
}

func (c *Mod_proxy) dirProxyStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyStatus directive
	return errors.New("ProxyStatus is not yet implemented")
}

func (c *Mod_proxy) dirProxyTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyTimeout directive
	return errors.New("ProxyTimeout is not yet implemented")
}

func (c *Mod_proxy) dirProxyVia(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyVia directive
	return errors.New("ProxyVia is not yet implemented")
}
