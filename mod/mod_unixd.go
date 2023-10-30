// https://httpd.apache.org/docs/2.4/mod/mod_unixd.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_unixd struct{}

func (c *Mod_unixd) DirChrootDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ChrootDir directive
	return errors.New("ChrootDir is not yet implemented")
}

func (c *Mod_unixd) DirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (c *Mod_unixd) DirSuexec(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Suexec directive
	return errors.New("Suexec is not yet implemented")
}

func (c *Mod_unixd) DirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
