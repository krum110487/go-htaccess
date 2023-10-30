// https://httpd.apache.org/docs/2.4/mod/mod_ldap.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_ldap struct{}

func (c *Mod_ldap) DirLDAPCacheEntries(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPCacheEntries directive
	return errors.New("LDAPCacheEntries is not yet implemented")
}

func (c *Mod_ldap) DirLDAPCacheTTL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPCacheTTL directive
	return errors.New("LDAPCacheTTL is not yet implemented")
}

func (c *Mod_ldap) DirLDAPConnectionPoolTTL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPConnectionPoolTTL directive
	return errors.New("LDAPConnectionPoolTTL is not yet implemented")
}

func (c *Mod_ldap) DirLDAPConnectionTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPConnectionTimeout directive
	return errors.New("LDAPConnectionTimeout is not yet implemented")
}

func (c *Mod_ldap) DirLDAPLibraryDebug(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPLibraryDebug directive
	return errors.New("LDAPLibraryDebug is not yet implemented")
}

func (c *Mod_ldap) DirLDAPOpCacheEntries(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPOpCacheEntries directive
	return errors.New("LDAPOpCacheEntries is not yet implemented")
}

func (c *Mod_ldap) DirLDAPOpCacheTTL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPOpCacheTTL directive
	return errors.New("LDAPOpCacheTTL is not yet implemented")
}

func (c *Mod_ldap) DirLDAPReferralHopLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPReferralHopLimit directive
	return errors.New("LDAPReferralHopLimit is not yet implemented")
}

func (c *Mod_ldap) DirLDAPReferrals(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPReferrals directive
	return errors.New("LDAPReferrals is not yet implemented")
}

func (c *Mod_ldap) DirLDAPRetries(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPRetries directive
	return errors.New("LDAPRetries is not yet implemented")
}

func (c *Mod_ldap) DirLDAPRetryDelay(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPRetryDelay directive
	return errors.New("LDAPRetryDelay is not yet implemented")
}

func (c *Mod_ldap) DirLDAPSharedCacheFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPSharedCacheFile directive
	return errors.New("LDAPSharedCacheFile is not yet implemented")
}

func (c *Mod_ldap) DirLDAPSharedCacheSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPSharedCacheSize directive
	return errors.New("LDAPSharedCacheSize is not yet implemented")
}

func (c *Mod_ldap) DirLDAPTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPTimeout directive
	return errors.New("LDAPTimeout is not yet implemented")
}

func (c *Mod_ldap) DirLDAPTrustedClientCert(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPTrustedClientCert directive
	return errors.New("LDAPTrustedClientCert is not yet implemented")
}

func (c *Mod_ldap) DirLDAPTrustedGlobalCert(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPTrustedGlobalCert directive
	return errors.New("LDAPTrustedGlobalCert is not yet implemented")
}

func (c *Mod_ldap) DirLDAPTrustedMode(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPTrustedMode directive
	return errors.New("LDAPTrustedMode is not yet implemented")
}

func (c *Mod_ldap) DirLDAPVerifyServerCert(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LDAPVerifyServerCert directive
	return errors.New("LDAPVerifyServerCert is not yet implemented")
}
