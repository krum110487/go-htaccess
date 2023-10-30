// https://httpd.apache.org/docs/2.4/mod/mod_alias.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_alias struct{}

func (c *Mod_alias) dirAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Alias directive
	return errors.New("Alias is not yet implemented")
}

func (c *Mod_alias) dirAliasMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AliasMatch directive
	return errors.New("AliasMatch is not yet implemented")
}

func (c *Mod_alias) dirAliasPreservePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AliasPreservePath directive
	return errors.New("AliasPreservePath is not yet implemented")
}

func (c *Mod_alias) dirRedirect(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Redirect directive
	return errors.New("Redirect is not yet implemented")
}

func (c *Mod_alias) dirRedirectMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectMatch directive
	return errors.New("RedirectMatch is not yet implemented")
}

func (c *Mod_alias) dirRedirectPermanent(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectPermanent directive
	return errors.New("RedirectPermanent is not yet implemented")
}

func (c *Mod_alias) dirRedirectRelative(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectRelative directive
	return errors.New("RedirectRelative is not yet implemented")
}

func (c *Mod_alias) dirRedirectTemp(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RedirectTemp directive
	return errors.New("RedirectTemp is not yet implemented")
}

func (c *Mod_alias) dirScriptAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptAlias directive
	return errors.New("ScriptAlias is not yet implemented")
}

func (c *Mod_alias) dirScriptAliasMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptAliasMatch directive
	return errors.New("ScriptAliasMatch is not yet implemented")
}
