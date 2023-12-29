// https://httpd.apache.org/docs/2.4/mod/mod_include.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_include struct{}

func (Mod_include) DirSSIEndTag(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSIEndTag directive
	return errors.New("SSIEndTag is not yet implemented")
}

func (Mod_include) DirSSIErrorMsg(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSIErrorMsg directive
	return errors.New("SSIErrorMsg is not yet implemented")
}

func (Mod_include) DirSSIETag(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSIETag directive
	return errors.New("SSIETag is not yet implemented")
}

func (Mod_include) DirSSILastModified(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSILastModified directive
	return errors.New("SSILastModified is not yet implemented")
}

func (Mod_include) DirSSILegacyExprParser(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSILegacyExprParser directive
	return errors.New("SSILegacyExprParser is not yet implemented")
}

func (Mod_include) DirSSIStartTag(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSIStartTag directive
	return errors.New("SSIStartTag is not yet implemented")
}

func (Mod_include) DirSSITimeFormat(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSITimeFormat directive
	return errors.New("SSITimeFormat is not yet implemented")
}

func (Mod_include) DirSSIUndefinedEcho(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SSIUndefinedEcho directive
	return errors.New("SSIUndefinedEcho is not yet implemented")
}

func (Mod_include) DirXBitHack(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: XBitHack directive
	return errors.New("XBitHack is not yet implemented")
}
