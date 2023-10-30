// https://httpd.apache.org/docs/2.4/mod/mod_proxy_html.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_proxy_html struct{}

func (c *Mod_proxy_html) DirProxyHTMLBufSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLBufSize directive
	return errors.New("ProxyHTMLBufSize is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLCharsetOut(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLCharsetOut directive
	return errors.New("ProxyHTMLCharsetOut is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLDocType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLDocType directive
	return errors.New("ProxyHTMLDocType is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLEnable directive
	return errors.New("ProxyHTMLEnable is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLEvents(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLEvents directive
	return errors.New("ProxyHTMLEvents is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLExtended(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLExtended directive
	return errors.New("ProxyHTMLExtended is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLFixups(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLFixups directive
	return errors.New("ProxyHTMLFixups is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLInterp(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLInterp directive
	return errors.New("ProxyHTMLInterp is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLLinks(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLLinks directive
	return errors.New("ProxyHTMLLinks is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLMeta(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLMeta directive
	return errors.New("ProxyHTMLMeta is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLStripComments(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLStripComments directive
	return errors.New("ProxyHTMLStripComments is not yet implemented")
}

func (c *Mod_proxy_html) DirProxyHTMLURLMap(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProxyHTMLURLMap directive
	return errors.New("ProxyHTMLURLMap is not yet implemented")
}
