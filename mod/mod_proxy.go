// https://httpd.apache.org/docs/2.4/mod/mod_proxy.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy struct{}

func (c *Mod_proxy) DirBalancerGrowth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerGrowth directive
	return errors.New("BalancerGrowth is not yet implemented")
}

func (c *Mod_proxy) DirBalancerInherit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerInherit directive
	return errors.New("BalancerInherit is not yet implemented")
}

func (c *Mod_proxy) DirBalancerMember(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerMember directive
	return errors.New("BalancerMember is not yet implemented")
}

func (c *Mod_proxy) DirBalancerPersist(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BalancerPersist directive
	return errors.New("BalancerPersist is not yet implemented")
}

func (c *Mod_proxy) DirNoProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: NoProxy directive
	return errors.New("NoProxy is not yet implemented")
}

func (c *Mod_proxy) DirProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Proxy> directive
	return errors.New("<Proxy> is not yet implemented")
}

func (c *Mod_proxy) DirProxy100Continue(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Proxy100Continue directive
	return errors.New("Proxy100Continue is not yet implemented")
}

func (c *Mod_proxy) DirProxyAddHeaders(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyAddHeaders directive
	return errors.New("ProxyAddHeaders is not yet implemented")
}

func (c *Mod_proxy) DirProxyBadHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyBadHeader directive
	return errors.New("ProxyBadHeader is not yet implemented")
}

func (c *Mod_proxy) DirProxyBlock(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyBlock directive
	return errors.New("ProxyBlock is not yet implemented")
}

func (c *Mod_proxy) DirProxyDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyDomain directive
	return errors.New("ProxyDomain is not yet implemented")
}

func (c *Mod_proxy) DirProxyErrorOverride(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyErrorOverride directive
	return errors.New("ProxyErrorOverride is not yet implemented")
}

func (c *Mod_proxy) DirProxyIOBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyIOBufferSize directive
	return errors.New("ProxyIOBufferSize is not yet implemented")
}

func (c *Mod_proxy) DirProxyMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <ProxyMatch> directive
	return errors.New("<ProxyMatch> is not yet implemented")
}

func (c *Mod_proxy) DirProxyMaxForwards(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyMaxForwards directive
	return errors.New("ProxyMaxForwards is not yet implemented")
}

func (c *Mod_proxy) DirProxyPass(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPass directive
	return errors.New("ProxyPass is not yet implemented")
}

func (c *Mod_proxy) DirProxyPassInherit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassInherit directive
	return errors.New("ProxyPassInherit is not yet implemented")
}

func (c *Mod_proxy) DirProxyPassInterpolateEnv(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassInterpolateEnv directive
	return errors.New("ProxyPassInterpolateEnv is not yet implemented")
}

func (c *Mod_proxy) DirProxyPassMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassMatch directive
	return errors.New("ProxyPassMatch is not yet implemented")
}

func (c *Mod_proxy) DirProxyPassReverse(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassReverse directive
	return errors.New("ProxyPassReverse is not yet implemented")
}

func (c *Mod_proxy) DirProxyPassReverseCookieDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassReverseCookieDomain directive
	return errors.New("ProxyPassReverseCookieDomain is not yet implemented")
}

func (c *Mod_proxy) DirProxyPassReverseCookiePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPassReverseCookiePath directive
	return errors.New("ProxyPassReverseCookiePath is not yet implemented")
}

func (c *Mod_proxy) DirProxyPreserveHost(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyPreserveHost directive
	return errors.New("ProxyPreserveHost is not yet implemented")
}

func (c *Mod_proxy) DirProxyReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyReceiveBufferSize directive
	return errors.New("ProxyReceiveBufferSize is not yet implemented")
}

func (c *Mod_proxy) DirProxyRemote(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyRemote directive
	return errors.New("ProxyRemote is not yet implemented")
}

func (c *Mod_proxy) DirProxyRemoteMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyRemoteMatch directive
	return errors.New("ProxyRemoteMatch is not yet implemented")
}

func (c *Mod_proxy) DirProxyRequests(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyRequests directive
	return errors.New("ProxyRequests is not yet implemented")
}

func (c *Mod_proxy) DirProxySet(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxySet directive
	return errors.New("ProxySet is not yet implemented")
}

func (c *Mod_proxy) DirProxySourceAddress(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxySourceAddress directive
	return errors.New("ProxySourceAddress is not yet implemented")
}

func (c *Mod_proxy) DirProxyStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyStatus directive
	return errors.New("ProxyStatus is not yet implemented")
}

func (c *Mod_proxy) DirProxyTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyTimeout directive
	return errors.New("ProxyTimeout is not yet implemented")
}

func (c *Mod_proxy) DirProxyVia(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyVia directive
	return errors.New("ProxyVia is not yet implemented")
}
