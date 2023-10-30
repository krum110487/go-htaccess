// https://httpd.apache.org/docs/2.4/mod/mod_dir.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dir struct{}

func (c *Mod_dir) dirDirectoryCheckHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DirectoryCheckHandler directive
	return errors.New("DirectoryCheckHandler is not yet implemented")
}

func (c *Mod_dir) dirDirectoryIndex(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DirectoryIndex directive
	return errors.New("DirectoryIndex is not yet implemented")
}

func (c *Mod_dir) dirDirectoryIndexRedirect(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DirectoryIndexRedirect directive
	return errors.New("DirectoryIndexRedirect is not yet implemented")
}

func (c *Mod_dir) dirDirectorySlash(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DirectorySlash directive
	return errors.New("DirectorySlash is not yet implemented")
}

func (c *Mod_dir) dirFallbackResource(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FallbackResource directive
	return errors.New("FallbackResource is not yet implemented")
}
