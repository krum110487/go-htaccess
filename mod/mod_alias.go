// https://httpd.apache.org/docs/2.4/mod/mod_alias.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_alias struct{}

func (c *Mod_alias) DirAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Alias directive
	return errors.New("Alias is not yet implemented")
}

func (c *Mod_alias) DirAliasMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AliasMatch directive
	return errors.New("AliasMatch is not yet implemented")
}

func (c *Mod_alias) DirAliasPreservePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AliasPreservePath directive
	return errors.New("AliasPreservePath is not yet implemented")
}

func (c *Mod_alias) DirRedirect(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Redirect directive
	return errors.New("Redirect is not yet implemented")
}

func (c *Mod_alias) DirRedirectMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectMatch directive
	return errors.New("RedirectMatch is not yet implemented")
}

func (c *Mod_alias) DirRedirectPermanent(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectPermanent directive
	return errors.New("RedirectPermanent is not yet implemented")
}

func (c *Mod_alias) DirRedirectRelative(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectRelative directive
	return errors.New("RedirectRelative is not yet implemented")
}

func (c *Mod_alias) DirRedirectTemp(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectTemp directive
	return errors.New("RedirectTemp is not yet implemented")
}

func (c *Mod_alias) DirScriptAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptAlias directive
	return errors.New("ScriptAlias is not yet implemented")
}

func (c *Mod_alias) DirScriptAliasMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptAliasMatch directive
	return errors.New("ScriptAliasMatch is not yet implemented")
}
