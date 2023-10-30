// https://httpd.apache.org/docs/2.4/mod/mod_session_dbd.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_session_dbd struct{}

func (c *Mod_session_dbd) dirSessionDBDCookieName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDCookieName directive
	return errors.New("SessionDBDCookieName is not yet implemented")
}

func (c *Mod_session_dbd) dirSessionDBDCookieName2(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDCookieName2 directive
	return errors.New("SessionDBDCookieName2 is not yet implemented")
}

func (c *Mod_session_dbd) dirSessionDBDCookieRemove(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDCookieRemove directive
	return errors.New("SessionDBDCookieRemove is not yet implemented")
}

func (c *Mod_session_dbd) dirSessionDBDDeleteLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDDeleteLabel directive
	return errors.New("SessionDBDDeleteLabel is not yet implemented")
}

func (c *Mod_session_dbd) dirSessionDBDInsertLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDInsertLabel directive
	return errors.New("SessionDBDInsertLabel is not yet implemented")
}

func (c *Mod_session_dbd) dirSessionDBDPerUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDPerUser directive
	return errors.New("SessionDBDPerUser is not yet implemented")
}

func (c *Mod_session_dbd) dirSessionDBDSelectLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDSelectLabel directive
	return errors.New("SessionDBDSelectLabel is not yet implemented")
}

func (c *Mod_session_dbd) dirSessionDBDUpdateLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionDBDUpdateLabel directive
	return errors.New("SessionDBDUpdateLabel is not yet implemented")
}
