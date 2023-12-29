// https://httpd.apache.org/docs/2.4/mod/mod_socache_memcache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_socache_memcache struct{}

func (Mod_socache_memcache) DirMemcacheConnTTL(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MemcacheConnTTL directive
	return errors.New("MemcacheConnTTL is not yet implemented")
}
