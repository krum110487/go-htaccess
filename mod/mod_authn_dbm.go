// https://httpd.apache.org/docs/2.4/mod/mod_authn_dbm.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authn_dbm struct{}

func (Mod_authn_dbm) DirAuthDBMType(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthDBMType directive
	return errors.New("AuthDBMType is not yet implemented")
}

func (Mod_authn_dbm) DirAuthDBMUserFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthDBMUserFile directive
	return errors.New("AuthDBMUserFile is not yet implemented")
}
