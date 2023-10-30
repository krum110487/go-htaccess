// https://httpd.apache.org/docs/2.4/mod/mod_include.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_include struct{}

func (c *Mod_include) dirSSIEndTag(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSIEndTag directive
	return errors.New("SSIEndTag is not yet implemented")
}

func (c *Mod_include) dirSSIErrorMsg(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSIErrorMsg directive
	return errors.New("SSIErrorMsg is not yet implemented")
}

func (c *Mod_include) dirSSIETag(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSIETag directive
	return errors.New("SSIETag is not yet implemented")
}

func (c *Mod_include) dirSSILastModified(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSILastModified directive
	return errors.New("SSILastModified is not yet implemented")
}

func (c *Mod_include) dirSSILegacyExprParser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSILegacyExprParser directive
	return errors.New("SSILegacyExprParser is not yet implemented")
}

func (c *Mod_include) dirSSIStartTag(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSIStartTag directive
	return errors.New("SSIStartTag is not yet implemented")
}

func (c *Mod_include) dirSSITimeFormat(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSITimeFormat directive
	return errors.New("SSITimeFormat is not yet implemented")
}

func (c *Mod_include) dirSSIUndefinedEcho(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSIUndefinedEcho directive
	return errors.New("SSIUndefinedEcho is not yet implemented")
}

func (c *Mod_include) dirXBitHack(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: XBitHack directive
	return errors.New("XBitHack is not yet implemented")
}
