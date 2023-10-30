// https://httpd.apache.org/docs/2.4/mod/mod_session.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_session struct{}

func (c *Mod_session) DirSession(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Session directive
	return errors.New("Session is not yet implemented")
}

func (c *Mod_session) DirSessionEnv(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionEnv directive
	return errors.New("SessionEnv is not yet implemented")
}

func (c *Mod_session) DirSessionExclude(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionExclude directive
	return errors.New("SessionExclude is not yet implemented")
}

func (c *Mod_session) DirSessionExpiryUpdateInterval(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionExpiryUpdateInterval directive
	return errors.New("SessionExpiryUpdateInterval is not yet implemented")
}

func (c *Mod_session) DirSessionHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionHeader directive
	return errors.New("SessionHeader is not yet implemented")
}

func (c *Mod_session) DirSessionInclude(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionInclude directive
	return errors.New("SessionInclude is not yet implemented")
}

func (c *Mod_session) DirSessionMaxAge(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionMaxAge directive
	return errors.New("SessionMaxAge is not yet implemented")
}
