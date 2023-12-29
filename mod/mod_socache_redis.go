// https://httpd.apache.org/docs/2.4/mod/mod_socache_redis.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_socache_redis struct{}

func (Mod_socache_redis) DirRedisConnPoolTTL(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RedisConnPoolTTL directive
	return errors.New("RedisConnPoolTTL is not yet implemented")
}

func (Mod_socache_redis) DirRedisTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RedisTimeout directive
	return errors.New("RedisTimeout is not yet implemented")
}
