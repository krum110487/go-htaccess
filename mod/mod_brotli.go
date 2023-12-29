// https://httpd.apache.org/docs/2.4/mod/mod_brotli.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_brotli struct{}

func (Mod_brotli) DirBrotliAlterETag(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: BrotliAlterETag directive
	return errors.New("BrotliAlterETag is not yet implemented")
}

func (Mod_brotli) DirBrotliCompressionMaxInputBlock(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: BrotliCompressionMaxInputBlock directive
	return errors.New("BrotliCompressionMaxInputBlock is not yet implemented")
}

func (Mod_brotli) DirBrotliCompressionQuality(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: BrotliCompressionQuality directive
	return errors.New("BrotliCompressionQuality is not yet implemented")
}

func (Mod_brotli) DirBrotliCompressionWindow(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: BrotliCompressionWindow directive
	return errors.New("BrotliCompressionWindow is not yet implemented")
}

func (Mod_brotli) DirBrotliFilterNote(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: BrotliFilterNote directive
	return errors.New("BrotliFilterNote is not yet implemented")
}
