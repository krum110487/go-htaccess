// https://httpd.apache.org/docs/2.4/mod/mod_file_cache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_file_cache struct{}

func (Mod_file_cache) DirCacheFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheFile directive
	return errors.New("CacheFile is not yet implemented")
}

func (Mod_file_cache) DirMMapFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MMapFile directive
	return errors.New("MMapFile is not yet implemented")
}
