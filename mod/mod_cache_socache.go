// https://httpd.apache.org/docs/2.4/mod/mod_cache_socache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cache_socache struct{}

func (c *Mod_cache_socache) dirCacheSocache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheSocache directive
	return errors.New("CacheSocache is not yet implemented")
}

func (c *Mod_cache_socache) dirCacheSocacheMaxSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheSocacheMaxSize directive
	return errors.New("CacheSocacheMaxSize is not yet implemented")
}

func (c *Mod_cache_socache) dirCacheSocacheMaxTime(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheSocacheMaxTime directive
	return errors.New("CacheSocacheMaxTime is not yet implemented")
}

func (c *Mod_cache_socache) dirCacheSocacheMinTime(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheSocacheMinTime directive
	return errors.New("CacheSocacheMinTime is not yet implemented")
}

func (c *Mod_cache_socache) dirCacheSocacheReadSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheSocacheReadSize directive
	return errors.New("CacheSocacheReadSize is not yet implemented")
}

func (c *Mod_cache_socache) dirCacheSocacheReadTime(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheSocacheReadTime directive
	return errors.New("CacheSocacheReadTime is not yet implemented")
}
