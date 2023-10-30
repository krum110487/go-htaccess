// https://httpd.apache.org/docs/2.4/mod/mod_socache_memcache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_socache_memcache struct{}

func (c *Mod_socache_memcache) dirMemcacheConnTTL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MemcacheConnTTL directive
	return errors.New("MemcacheConnTTL is not yet implemented")
}
