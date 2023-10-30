// https://httpd.apache.org/docs/2.4/mod/mod_cache_disk.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cache_disk struct{}

func (c *Mod_cache_disk) dirCacheDirLength(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheDirLength directive
	return errors.New("CacheDirLength is not yet implemented")
}

func (c *Mod_cache_disk) dirCacheDirLevels(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheDirLevels directive
	return errors.New("CacheDirLevels is not yet implemented")
}

func (c *Mod_cache_disk) dirCacheMaxFileSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheMaxFileSize directive
	return errors.New("CacheMaxFileSize is not yet implemented")
}

func (c *Mod_cache_disk) dirCacheMinFileSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheMinFileSize directive
	return errors.New("CacheMinFileSize is not yet implemented")
}

func (c *Mod_cache_disk) dirCacheReadSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheReadSize directive
	return errors.New("CacheReadSize is not yet implemented")
}

func (c *Mod_cache_disk) dirCacheReadTime(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheReadTime directive
	return errors.New("CacheReadTime is not yet implemented")
}

func (c *Mod_cache_disk) dirCacheRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheRoot directive
	return errors.New("CacheRoot is not yet implemented")
}
