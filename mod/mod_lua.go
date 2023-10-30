// https://httpd.apache.org/docs/2.4/mod/mod_lua.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_lua struct{}

func (c *Mod_lua) dirLuaAuthzProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaAuthzProvider directive
	return errors.New("LuaAuthzProvider is not yet implemented")
}

func (c *Mod_lua) dirLuaCodeCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaCodeCache directive
	return errors.New("LuaCodeCache is not yet implemented")
}

func (c *Mod_lua) dirLuaHookAccessChecker(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookAccessChecker directive
	return errors.New("LuaHookAccessChecker is not yet implemented")
}

func (c *Mod_lua) dirLuaHookAuthChecker(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookAuthChecker directive
	return errors.New("LuaHookAuthChecker is not yet implemented")
}

func (c *Mod_lua) dirLuaHookCheckUserID(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookCheckUserID directive
	return errors.New("LuaHookCheckUserID is not yet implemented")
}

func (c *Mod_lua) dirLuaHookFixups(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookFixups directive
	return errors.New("LuaHookFixups is not yet implemented")
}

func (c *Mod_lua) dirLuaHookInsertFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookInsertFilter directive
	return errors.New("LuaHookInsertFilter is not yet implemented")
}

func (c *Mod_lua) dirLuaHookLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookLog directive
	return errors.New("LuaHookLog is not yet implemented")
}

func (c *Mod_lua) dirLuaHookMapToStorage(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookMapToStorage directive
	return errors.New("LuaHookMapToStorage is not yet implemented")
}

func (c *Mod_lua) dirLuaHookPreTranslate(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookPreTranslate directive
	return errors.New("LuaHookPreTranslate is not yet implemented")
}

func (c *Mod_lua) dirLuaHookTranslateName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookTranslateName directive
	return errors.New("LuaHookTranslateName is not yet implemented")
}

func (c *Mod_lua) dirLuaHookTypeChecker(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaHookTypeChecker directive
	return errors.New("LuaHookTypeChecker is not yet implemented")
}

func (c *Mod_lua) dirLuaInherit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaInherit directive
	return errors.New("LuaInherit is not yet implemented")
}

func (c *Mod_lua) dirLuaInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaInputFilter directive
	return errors.New("LuaInputFilter is not yet implemented")
}

func (c *Mod_lua) dirLuaMapHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaMapHandler directive
	return errors.New("LuaMapHandler is not yet implemented")
}

func (c *Mod_lua) dirLuaOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaOutputFilter directive
	return errors.New("LuaOutputFilter is not yet implemented")
}

func (c *Mod_lua) dirLuaPackageCPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaPackageCPath directive
	return errors.New("LuaPackageCPath is not yet implemented")
}

func (c *Mod_lua) dirLuaPackagePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaPackagePath directive
	return errors.New("LuaPackagePath is not yet implemented")
}

func (c *Mod_lua) dirLuaQuickHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaQuickHandler directive
	return errors.New("LuaQuickHandler is not yet implemented")
}

func (c *Mod_lua) dirLuaRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaRoot directive
	return errors.New("LuaRoot is not yet implemented")
}

func (c *Mod_lua) dirLuaScope(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LuaScope directive
	return errors.New("LuaScope is not yet implemented")
}
