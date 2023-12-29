// https://httpd.apache.org/docs/2.4/mod/mod_authn_anon.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authn_anon struct{}

func (Mod_authn_anon) DirAnonymous(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Anonymous directive
	return errors.New("Anonymous is not yet implemented")
}

func (Mod_authn_anon) DirAnonymous_LogEmail(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Anonymous_LogEmail directive
	return errors.New("Anonymous_LogEmail is not yet implemented")
}

func (Mod_authn_anon) DirAnonymous_MustGiveEmail(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Anonymous_MustGiveEmail directive
	return errors.New("Anonymous_MustGiveEmail is not yet implemented")
}

func (Mod_authn_anon) DirAnonymous_NoUserID(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Anonymous_NoUserID directive
	return errors.New("Anonymous_NoUserID is not yet implemented")
}

func (Mod_authn_anon) DirAnonymous_VerifyEmail(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Anonymous_VerifyEmail directive
	return errors.New("Anonymous_VerifyEmail is not yet implemented")
}
