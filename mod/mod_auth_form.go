// https://httpd.apache.org/docs/2.4/mod/mod_auth_form.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_auth_form struct{}

func (c *Mod_auth_form) dirAuthFormAuthoritative(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormAuthoritative directive
	return errors.New("AuthFormAuthoritative is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormBody(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormBody directive
	return errors.New("AuthFormBody is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormDisableNoStore(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormDisableNoStore directive
	return errors.New("AuthFormDisableNoStore is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormFakeBasicAuth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormFakeBasicAuth directive
	return errors.New("AuthFormFakeBasicAuth is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormLocation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormLocation directive
	return errors.New("AuthFormLocation is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormLoginRequiredLocation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormLoginRequiredLocation directive
	return errors.New("AuthFormLoginRequiredLocation is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormLoginSuccessLocation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormLoginSuccessLocation directive
	return errors.New("AuthFormLoginSuccessLocation is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormLogoutLocation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormLogoutLocation directive
	return errors.New("AuthFormLogoutLocation is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormMethod(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormMethod directive
	return errors.New("AuthFormMethod is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormMimetype(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormMimetype directive
	return errors.New("AuthFormMimetype is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormPassword(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormPassword directive
	return errors.New("AuthFormPassword is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormProvider directive
	return errors.New("AuthFormProvider is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormSitePassphrase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormSitePassphrase directive
	return errors.New("AuthFormSitePassphrase is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormSize directive
	return errors.New("AuthFormSize is not yet implemented")
}

func (c *Mod_auth_form) dirAuthFormUsername(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthFormUsername directive
	return errors.New("AuthFormUsername is not yet implemented")
}
