// https://httpd.apache.org/docs/2.4/mod/mod_session_dbd.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_session_dbd struct{}

func (Mod_session_dbd) DirSessionDBDCookieName(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDCookieName directive
	return errors.New("SessionDBDCookieName is not yet implemented")
}

func (Mod_session_dbd) DirSessionDBDCookieName2(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDCookieName2 directive
	return errors.New("SessionDBDCookieName2 is not yet implemented")
}

func (Mod_session_dbd) DirSessionDBDCookieRemove(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDCookieRemove directive
	return errors.New("SessionDBDCookieRemove is not yet implemented")
}

func (Mod_session_dbd) DirSessionDBDDeleteLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDDeleteLabel directive
	return errors.New("SessionDBDDeleteLabel is not yet implemented")
}

func (Mod_session_dbd) DirSessionDBDInsertLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDInsertLabel directive
	return errors.New("SessionDBDInsertLabel is not yet implemented")
}

func (Mod_session_dbd) DirSessionDBDPerUser(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDPerUser directive
	return errors.New("SessionDBDPerUser is not yet implemented")
}

func (Mod_session_dbd) DirSessionDBDSelectLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDSelectLabel directive
	return errors.New("SessionDBDSelectLabel is not yet implemented")
}

func (Mod_session_dbd) DirSessionDBDUpdateLabel(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionDBDUpdateLabel directive
	return errors.New("SessionDBDUpdateLabel is not yet implemented")
}
