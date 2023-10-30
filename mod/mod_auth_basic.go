// https://httpd.apache.org/docs/2.4/mod/mod_auth_basic.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_auth_basic struct{}

func (c *Mod_auth_basic) DirAuthBasicAuthoritative(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthBasicAuthoritative directive
	return errors.New("AuthBasicAuthoritative is not yet implemented")
}

func (c *Mod_auth_basic) DirAuthBasicFake(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthBasicFake directive
	return errors.New("AuthBasicFake is not yet implemented")
}

func (c *Mod_auth_basic) DirAuthBasicProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthBasicProvider directive
	return errors.New("AuthBasicProvider is not yet implemented")
}

func (c *Mod_auth_basic) DirAuthBasicUseDigestAlgorithm(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthBasicUseDigestAlgorithm directive
	return errors.New("AuthBasicUseDigestAlgorithm is not yet implemented")
}
