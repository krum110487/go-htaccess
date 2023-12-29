// https://httpd.apache.org/docs/2.4/mod/mod_cache_socache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cache_socache struct{}

func (Mod_cache_socache) DirCacheSocache(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheSocache directive
	return errors.New("CacheSocache is not yet implemented")
}

func (Mod_cache_socache) DirCacheSocacheMaxSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheSocacheMaxSize directive
	return errors.New("CacheSocacheMaxSize is not yet implemented")
}

func (Mod_cache_socache) DirCacheSocacheMaxTime(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheSocacheMaxTime directive
	return errors.New("CacheSocacheMaxTime is not yet implemented")
}

func (Mod_cache_socache) DirCacheSocacheMinTime(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheSocacheMinTime directive
	return errors.New("CacheSocacheMinTime is not yet implemented")
}

func (Mod_cache_socache) DirCacheSocacheReadSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheSocacheReadSize directive
	return errors.New("CacheSocacheReadSize is not yet implemented")
}

func (Mod_cache_socache) DirCacheSocacheReadTime(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheSocacheReadTime directive
	return errors.New("CacheSocacheReadTime is not yet implemented")
}
