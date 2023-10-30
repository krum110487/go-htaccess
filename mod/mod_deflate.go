// https://httpd.apache.org/docs/2.4/mod/mod_deflate.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_deflate struct{}

func (c *Mod_deflate) dirDeflateAlterETag(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateAlterETag directive
	return errors.New("DeflateAlterETag is not yet implemented")
}

func (c *Mod_deflate) dirDeflateBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateBufferSize directive
	return errors.New("DeflateBufferSize is not yet implemented")
}

func (c *Mod_deflate) dirDeflateCompressionLevel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateCompressionLevel directive
	return errors.New("DeflateCompressionLevel is not yet implemented")
}

func (c *Mod_deflate) dirDeflateFilterNote(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateFilterNote directive
	return errors.New("DeflateFilterNote is not yet implemented")
}

func (c *Mod_deflate) dirDeflateInflateLimitRequestBody(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateInflateLimitRequestBody directive
	return errors.New("DeflateInflateLimitRequestBody is not yet implemented")
}

func (c *Mod_deflate) dirDeflateInflateRatioBurst(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateInflateRatioBurst directive
	return errors.New("DeflateInflateRatioBurst is not yet implemented")
}

func (c *Mod_deflate) dirDeflateInflateRatioLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateInflateRatioLimit directive
	return errors.New("DeflateInflateRatioLimit is not yet implemented")
}

func (c *Mod_deflate) dirDeflateMemLevel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateMemLevel directive
	return errors.New("DeflateMemLevel is not yet implemented")
}

func (c *Mod_deflate) dirDeflateWindowSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DeflateWindowSize directive
	return errors.New("DeflateWindowSize is not yet implemented")
}
