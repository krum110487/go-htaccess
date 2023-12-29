// https://httpd.apache.org/docs/2.4/mod/mod_cache_disk.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cache_disk struct{}

func (Mod_cache_disk) DirCacheDirLength(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheDirLength directive
	return errors.New("CacheDirLength is not yet implemented")
}

func (Mod_cache_disk) DirCacheDirLevels(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheDirLevels directive
	return errors.New("CacheDirLevels is not yet implemented")
}

func (Mod_cache_disk) DirCacheMaxFileSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheMaxFileSize directive
	return errors.New("CacheMaxFileSize is not yet implemented")
}

func (Mod_cache_disk) DirCacheMinFileSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheMinFileSize directive
	return errors.New("CacheMinFileSize is not yet implemented")
}

func (Mod_cache_disk) DirCacheReadSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheReadSize directive
	return errors.New("CacheReadSize is not yet implemented")
}

func (Mod_cache_disk) DirCacheReadTime(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheReadTime directive
	return errors.New("CacheReadTime is not yet implemented")
}

func (Mod_cache_disk) DirCacheRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheRoot directive
	return errors.New("CacheRoot is not yet implemented")
}
