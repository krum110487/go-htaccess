// https://httpd.apache.org/docs/2.4/mod/mod_isapi.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_isapi struct{}

func (c *Mod_isapi) DirISAPIAppendLogToErrors(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIAppendLogToErrors directive
	return errors.New("ISAPIAppendLogToErrors is not yet implemented")
}

func (c *Mod_isapi) DirISAPIAppendLogToQuery(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIAppendLogToQuery directive
	return errors.New("ISAPIAppendLogToQuery is not yet implemented")
}

func (c *Mod_isapi) DirISAPICacheFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPICacheFile directive
	return errors.New("ISAPICacheFile is not yet implemented")
}

func (c *Mod_isapi) DirISAPIFakeAsync(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIFakeAsync directive
	return errors.New("ISAPIFakeAsync is not yet implemented")
}

func (c *Mod_isapi) DirISAPILogNotSupported(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPILogNotSupported directive
	return errors.New("ISAPILogNotSupported is not yet implemented")
}

func (c *Mod_isapi) DirISAPIReadAheadBuffer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIReadAheadBuffer directive
	return errors.New("ISAPIReadAheadBuffer is not yet implemented")
}
