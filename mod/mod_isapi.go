// https://httpd.apache.org/docs/2.4/mod/mod_isapi.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_isapi struct{}

func (c *Mod_isapi) dirISAPIAppendLogToErrors(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIAppendLogToErrors directive
	return errors.New("ISAPIAppendLogToErrors is not yet implemented")
}

func (c *Mod_isapi) dirISAPIAppendLogToQuery(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIAppendLogToQuery directive
	return errors.New("ISAPIAppendLogToQuery is not yet implemented")
}

func (c *Mod_isapi) dirISAPICacheFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPICacheFile directive
	return errors.New("ISAPICacheFile is not yet implemented")
}

func (c *Mod_isapi) dirISAPIFakeAsync(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIFakeAsync directive
	return errors.New("ISAPIFakeAsync is not yet implemented")
}

func (c *Mod_isapi) dirISAPILogNotSupported(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPILogNotSupported directive
	return errors.New("ISAPILogNotSupported is not yet implemented")
}

func (c *Mod_isapi) dirISAPIReadAheadBuffer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ISAPIReadAheadBuffer directive
	return errors.New("ISAPIReadAheadBuffer is not yet implemented")
}
