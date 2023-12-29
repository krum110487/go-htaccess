// https://httpd.apache.org/docs/2.4/mod/mod_privileges.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_privileges struct{}

func (Mod_privileges) DirDTracePrivileges(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DTracePrivileges directive
	return errors.New("DTracePrivileges is not yet implemented")
}

func (Mod_privileges) DirPrivilegesMode(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: PrivilegesMode directive
	return errors.New("PrivilegesMode is not yet implemented")
}

func (Mod_privileges) DirVHostCGIMode(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: VHostCGIMode directive
	return errors.New("VHostCGIMode is not yet implemented")
}

func (Mod_privileges) DirVHostCGIPrivs(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: VHostCGIPrivs directive
	return errors.New("VHostCGIPrivs is not yet implemented")
}

func (Mod_privileges) DirVHostGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: VHostGroup directive
	return errors.New("VHostGroup is not yet implemented")
}

func (Mod_privileges) DirVHostPrivs(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: VHostPrivs directive
	return errors.New("VHostPrivs is not yet implemented")
}

func (Mod_privileges) DirVHostSecure(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: VHostSecure directive
	return errors.New("VHostSecure is not yet implemented")
}

func (Mod_privileges) DirVHostUser(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: VHostUser directive
	return errors.New("VHostUser is not yet implemented")
}
