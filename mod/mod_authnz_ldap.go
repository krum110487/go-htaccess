// https://httpd.apache.org/docs/2.4/mod/mod_authnz_ldap.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_authnz_ldap struct{}

func (c *Mod_authnz_ldap) dirAuthLDAPAuthorizePrefix(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPAuthorizePrefix directive
	return errors.New("AuthLDAPAuthorizePrefix is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPBindAuthoritative(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPBindAuthoritative directive
	return errors.New("AuthLDAPBindAuthoritative is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPBindDN(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPBindDN directive
	return errors.New("AuthLDAPBindDN is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPBindPassword(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPBindPassword directive
	return errors.New("AuthLDAPBindPassword is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPCharsetConfig(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPCharsetConfig directive
	return errors.New("AuthLDAPCharsetConfig is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPCompareAsUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPCompareAsUser directive
	return errors.New("AuthLDAPCompareAsUser is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPCompareDNOnServer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPCompareDNOnServer directive
	return errors.New("AuthLDAPCompareDNOnServer is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPDereferenceAliases(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPDereferenceAliases directive
	return errors.New("AuthLDAPDereferenceAliases is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPGroupAttribute(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPGroupAttribute directive
	return errors.New("AuthLDAPGroupAttribute is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPGroupAttributeIsDN(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPGroupAttributeIsDN directive
	return errors.New("AuthLDAPGroupAttributeIsDN is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPInitialBindAsUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPInitialBindAsUser directive
	return errors.New("AuthLDAPInitialBindAsUser is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPInitialBindPattern(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPInitialBindPattern directive
	return errors.New("AuthLDAPInitialBindPattern is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPMaxSubGroupDepth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPMaxSubGroupDepth directive
	return errors.New("AuthLDAPMaxSubGroupDepth is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPRemoteUserAttribute(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPRemoteUserAttribute directive
	return errors.New("AuthLDAPRemoteUserAttribute is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPRemoteUserIsDN(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPRemoteUserIsDN directive
	return errors.New("AuthLDAPRemoteUserIsDN is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPSearchAsUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPSearchAsUser directive
	return errors.New("AuthLDAPSearchAsUser is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPSubGroupAttribute(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPSubGroupAttribute directive
	return errors.New("AuthLDAPSubGroupAttribute is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPSubGroupClass(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPSubGroupClass directive
	return errors.New("AuthLDAPSubGroupClass is not yet implemented")
}

func (c *Mod_authnz_ldap) dirAuthLDAPURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthLDAPURL directive
	return errors.New("AuthLDAPURL is not yet implemented")
}
