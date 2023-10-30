// https://httpd.apache.org/docs/2.4/mod/mod_dbd.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dbd struct{}

func (c *Mod_dbd) DirDBDExptime(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDExptime directive
	return errors.New("DBDExptime is not yet implemented")
}

func (c *Mod_dbd) DirDBDInitSQL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDInitSQL directive
	return errors.New("DBDInitSQL is not yet implemented")
}

func (c *Mod_dbd) DirDBDKeep(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDKeep directive
	return errors.New("DBDKeep is not yet implemented")
}

func (c *Mod_dbd) DirDBDMax(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDMax directive
	return errors.New("DBDMax is not yet implemented")
}

func (c *Mod_dbd) DirDBDMin(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDMin directive
	return errors.New("DBDMin is not yet implemented")
}

func (c *Mod_dbd) DirDBDParams(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDParams directive
	return errors.New("DBDParams is not yet implemented")
}

func (c *Mod_dbd) DirDBDPersist(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDPersist directive
	return errors.New("DBDPersist is not yet implemented")
}

func (c *Mod_dbd) DirDBDPrepareSQL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDPrepareSQL directive
	return errors.New("DBDPrepareSQL is not yet implemented")
}

func (c *Mod_dbd) DirDBDriver(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DBDriver directive
	return errors.New("DBDriver is not yet implemented")
}
