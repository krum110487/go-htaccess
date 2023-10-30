// https://httpd.apache.org/docs/2.4/mod/mod_cache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cache struct{}

func (c *Mod_cache) dirCacheDefaultExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheDefaultExpire directive
	return errors.New("CacheDefaultExpire is not yet implemented")
}

func (c *Mod_cache) dirCacheDetailHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheDetailHeader directive
	return errors.New("CacheDetailHeader is not yet implemented")
}

func (c *Mod_cache) dirCacheDisable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheDisable directive
	return errors.New("CacheDisable is not yet implemented")
}

func (c *Mod_cache) dirCacheEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheEnable directive
	return errors.New("CacheEnable is not yet implemented")
}

func (c *Mod_cache) dirCacheHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheHeader directive
	return errors.New("CacheHeader is not yet implemented")
}

func (c *Mod_cache) dirCacheIgnoreCacheControl(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheIgnoreCacheControl directive
	return errors.New("CacheIgnoreCacheControl is not yet implemented")
}

func (c *Mod_cache) dirCacheIgnoreHeaders(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheIgnoreHeaders directive
	return errors.New("CacheIgnoreHeaders is not yet implemented")
}

func (c *Mod_cache) dirCacheIgnoreNoLastMod(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheIgnoreNoLastMod directive
	return errors.New("CacheIgnoreNoLastMod is not yet implemented")
}

func (c *Mod_cache) dirCacheIgnoreQueryString(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheIgnoreQueryString directive
	return errors.New("CacheIgnoreQueryString is not yet implemented")
}

func (c *Mod_cache) dirCacheIgnoreURLSessionIdentifiers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheIgnoreURLSessionIdentifiers directive
	return errors.New("CacheIgnoreURLSessionIdentifiers is not yet implemented")
}

func (c *Mod_cache) dirCacheKeyBaseURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheKeyBaseURL directive
	return errors.New("CacheKeyBaseURL is not yet implemented")
}

func (c *Mod_cache) dirCacheLastModifiedFactor(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheLastModifiedFactor directive
	return errors.New("CacheLastModifiedFactor is not yet implemented")
}

func (c *Mod_cache) dirCacheLock(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheLock directive
	return errors.New("CacheLock is not yet implemented")
}

func (c *Mod_cache) dirCacheLockMaxAge(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheLockMaxAge directive
	return errors.New("CacheLockMaxAge is not yet implemented")
}

func (c *Mod_cache) dirCacheLockPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheLockPath directive
	return errors.New("CacheLockPath is not yet implemented")
}

func (c *Mod_cache) dirCacheMaxExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheMaxExpire directive
	return errors.New("CacheMaxExpire is not yet implemented")
}

func (c *Mod_cache) dirCacheMinExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheMinExpire directive
	return errors.New("CacheMinExpire is not yet implemented")
}

func (c *Mod_cache) dirCacheQuickHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheQuickHandler directive
	return errors.New("CacheQuickHandler is not yet implemented")
}

func (c *Mod_cache) dirCacheStaleOnError(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheStaleOnError directive
	return errors.New("CacheStaleOnError is not yet implemented")
}

func (c *Mod_cache) dirCacheStoreExpired(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheStoreExpired directive
	return errors.New("CacheStoreExpired is not yet implemented")
}

func (c *Mod_cache) dirCacheStoreNoStore(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheStoreNoStore directive
	return errors.New("CacheStoreNoStore is not yet implemented")
}

func (c *Mod_cache) dirCacheStorePrivate(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheStorePrivate directive
	return errors.New("CacheStorePrivate is not yet implemented")
}
