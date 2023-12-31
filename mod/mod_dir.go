// https://httpd.apache.org/docs/2.4/mod/mod_dir.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_dir struct{}

func (Mod_dir) DirDirectoryCheckHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DirectoryCheckHandler directive
	return errors.New("DirectoryCheckHandler is not yet implemented")
}

func (Mod_dir) DirDirectoryIndex(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DirectoryIndex directive
	return errors.New("DirectoryIndex is not yet implemented")
}

func (Mod_dir) DirDirectoryIndexRedirect(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DirectoryIndexRedirect directive
	return errors.New("DirectoryIndexRedirect is not yet implemented")
}

func (Mod_dir) DirDirectorySlash(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DirectorySlash directive
	return errors.New("DirectorySlash is not yet implemented")
}

func (Mod_dir) DirFallbackResource(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: FallbackResource directive
	return errors.New("FallbackResource is not yet implemented")
}
