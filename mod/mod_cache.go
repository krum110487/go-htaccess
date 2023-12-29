// https://httpd.apache.org/docs/2.4/mod/mod_cache.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cache struct{}

func (Mod_cache) DirCacheDefaultExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheDefaultExpire directive
	return errors.New("CacheDefaultExpire is not yet implemented")
}

func (Mod_cache) DirCacheDetailHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheDetailHeader directive
	return errors.New("CacheDetailHeader is not yet implemented")
}

func (Mod_cache) DirCacheDisable(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheDisable directive
	return errors.New("CacheDisable is not yet implemented")
}

func (Mod_cache) DirCacheEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheEnable directive
	return errors.New("CacheEnable is not yet implemented")
}

func (Mod_cache) DirCacheHeader(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheHeader directive
	return errors.New("CacheHeader is not yet implemented")
}

func (Mod_cache) DirCacheIgnoreCacheControl(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheIgnoreCacheControl directive
	return errors.New("CacheIgnoreCacheControl is not yet implemented")
}

func (Mod_cache) DirCacheIgnoreHeaders(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheIgnoreHeaders directive
	return errors.New("CacheIgnoreHeaders is not yet implemented")
}

func (Mod_cache) DirCacheIgnoreNoLastMod(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheIgnoreNoLastMod directive
	return errors.New("CacheIgnoreNoLastMod is not yet implemented")
}

func (Mod_cache) DirCacheIgnoreQueryString(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheIgnoreQueryString directive
	return errors.New("CacheIgnoreQueryString is not yet implemented")
}

func (Mod_cache) DirCacheIgnoreURLSessionIdentifiers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheIgnoreURLSessionIdentifiers directive
	return errors.New("CacheIgnoreURLSessionIdentifiers is not yet implemented")
}

func (Mod_cache) DirCacheKeyBaseURL(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheKeyBaseURL directive
	return errors.New("CacheKeyBaseURL is not yet implemented")
}

func (Mod_cache) DirCacheLastModifiedFactor(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheLastModifiedFactor directive
	return errors.New("CacheLastModifiedFactor is not yet implemented")
}

func (Mod_cache) DirCacheLock(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheLock directive
	return errors.New("CacheLock is not yet implemented")
}

func (Mod_cache) DirCacheLockMaxAge(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheLockMaxAge directive
	return errors.New("CacheLockMaxAge is not yet implemented")
}

func (Mod_cache) DirCacheLockPath(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheLockPath directive
	return errors.New("CacheLockPath is not yet implemented")
}

func (Mod_cache) DirCacheMaxExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheMaxExpire directive
	return errors.New("CacheMaxExpire is not yet implemented")
}

func (Mod_cache) DirCacheMinExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheMinExpire directive
	return errors.New("CacheMinExpire is not yet implemented")
}

func (Mod_cache) DirCacheQuickHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheQuickHandler directive
	return errors.New("CacheQuickHandler is not yet implemented")
}

func (Mod_cache) DirCacheStaleOnError(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheStaleOnError directive
	return errors.New("CacheStaleOnError is not yet implemented")
}

func (Mod_cache) DirCacheStoreExpired(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheStoreExpired directive
	return errors.New("CacheStoreExpired is not yet implemented")
}

func (Mod_cache) DirCacheStoreNoStore(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheStoreNoStore directive
	return errors.New("CacheStoreNoStore is not yet implemented")
}

func (Mod_cache) DirCacheStorePrivate(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CacheStorePrivate directive
	return errors.New("CacheStorePrivate is not yet implemented")
}
