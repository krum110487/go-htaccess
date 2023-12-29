// https://httpd.apache.org/docs/2.4/mod/mod_authn_dbd.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authn_dbd struct{}

func (Mod_authn_dbd) DirAuthDBDUserPWQuery(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthDBDUserPWQuery directive
	return errors.New("AuthDBDUserPWQuery is not yet implemented")
}

func (Mod_authn_dbd) DirAuthDBDUserRealmQuery(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AuthDBDUserRealmQuery directive
	return errors.New("AuthDBDUserRealmQuery is not yet implemented")
}
