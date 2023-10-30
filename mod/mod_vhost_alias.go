// https://httpd.apache.org/docs/2.4/mod/mod_vhost_alias.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_vhost_alias struct{}

func (c *Mod_vhost_alias) dirVirtualDocumentRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VirtualDocumentRoot directive
	return errors.New("VirtualDocumentRoot is not yet implemented")
}

func (c *Mod_vhost_alias) dirVirtualDocumentRootIP(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VirtualDocumentRootIP directive
	return errors.New("VirtualDocumentRootIP is not yet implemented")
}

func (c *Mod_vhost_alias) dirVirtualScriptAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VirtualScriptAlias directive
	return errors.New("VirtualScriptAlias is not yet implemented")
}

func (c *Mod_vhost_alias) dirVirtualScriptAliasIP(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VirtualScriptAliasIP directive
	return errors.New("VirtualScriptAliasIP is not yet implemented")
}
