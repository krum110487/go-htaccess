// https://httpd.apache.org/docs/2.4/mod/mod_proxy_html.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_html struct{}

func (c *Mod_proxy_html) dirProxyHTMLBufSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLBufSize directive
	return errors.New("ProxyHTMLBufSize is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLCharsetOut(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLCharsetOut directive
	return errors.New("ProxyHTMLCharsetOut is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLDocType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLDocType directive
	return errors.New("ProxyHTMLDocType is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLEnable directive
	return errors.New("ProxyHTMLEnable is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLEvents(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLEvents directive
	return errors.New("ProxyHTMLEvents is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLExtended(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLExtended directive
	return errors.New("ProxyHTMLExtended is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLFixups(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLFixups directive
	return errors.New("ProxyHTMLFixups is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLInterp(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLInterp directive
	return errors.New("ProxyHTMLInterp is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLLinks(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLLinks directive
	return errors.New("ProxyHTMLLinks is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLMeta(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLMeta directive
	return errors.New("ProxyHTMLMeta is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLStripComments(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLStripComments directive
	return errors.New("ProxyHTMLStripComments is not yet implemented")
}

func (c *Mod_proxy_html) dirProxyHTMLURLMap(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLURLMap directive
	return errors.New("ProxyHTMLURLMap is not yet implemented")
}
