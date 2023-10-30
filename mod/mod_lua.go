// https://httpd.apache.org/docs/2.4/mod/mod_lua.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_lua struct{}

func (c *Mod_lua) DirLuaAuthzProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaAuthzProvider directive
	return errors.New("LuaAuthzProvider is not yet implemented")
}

func (c *Mod_lua) DirLuaCodeCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaCodeCache directive
	return errors.New("LuaCodeCache is not yet implemented")
}

func (c *Mod_lua) DirLuaHookAccessChecker(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookAccessChecker directive
	return errors.New("LuaHookAccessChecker is not yet implemented")
}

func (c *Mod_lua) DirLuaHookAuthChecker(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookAuthChecker directive
	return errors.New("LuaHookAuthChecker is not yet implemented")
}

func (c *Mod_lua) DirLuaHookCheckUserID(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookCheckUserID directive
	return errors.New("LuaHookCheckUserID is not yet implemented")
}

func (c *Mod_lua) DirLuaHookFixups(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookFixups directive
	return errors.New("LuaHookFixups is not yet implemented")
}

func (c *Mod_lua) DirLuaHookInsertFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookInsertFilter directive
	return errors.New("LuaHookInsertFilter is not yet implemented")
}

func (c *Mod_lua) DirLuaHookLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookLog directive
	return errors.New("LuaHookLog is not yet implemented")
}

func (c *Mod_lua) DirLuaHookMapToStorage(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookMapToStorage directive
	return errors.New("LuaHookMapToStorage is not yet implemented")
}

func (c *Mod_lua) DirLuaHookPreTranslate(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookPreTranslate directive
	return errors.New("LuaHookPreTranslate is not yet implemented")
}

func (c *Mod_lua) DirLuaHookTranslateName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookTranslateName directive
	return errors.New("LuaHookTranslateName is not yet implemented")
}

func (c *Mod_lua) DirLuaHookTypeChecker(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookTypeChecker directive
	return errors.New("LuaHookTypeChecker is not yet implemented")
}

func (c *Mod_lua) DirLuaInherit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaInherit directive
	return errors.New("LuaInherit is not yet implemented")
}

func (c *Mod_lua) DirLuaInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaInputFilter directive
	return errors.New("LuaInputFilter is not yet implemented")
}

func (c *Mod_lua) DirLuaMapHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaMapHandler directive
	return errors.New("LuaMapHandler is not yet implemented")
}

func (c *Mod_lua) DirLuaOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaOutputFilter directive
	return errors.New("LuaOutputFilter is not yet implemented")
}

func (c *Mod_lua) DirLuaPackageCPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaPackageCPath directive
	return errors.New("LuaPackageCPath is not yet implemented")
}

func (c *Mod_lua) DirLuaPackagePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaPackagePath directive
	return errors.New("LuaPackagePath is not yet implemented")
}

func (c *Mod_lua) DirLuaQuickHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaQuickHandler directive
	return errors.New("LuaQuickHandler is not yet implemented")
}

func (c *Mod_lua) DirLuaRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaRoot directive
	return errors.New("LuaRoot is not yet implemented")
}

func (c *Mod_lua) DirLuaScope(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaScope directive
	return errors.New("LuaScope is not yet implemented")
}
