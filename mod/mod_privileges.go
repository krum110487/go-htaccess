// https://httpd.apache.org/docs/2.4/mod/mod_privileges.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_privileges struct{}

func (c *Mod_privileges) dirDTracePrivileges(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DTracePrivileges directive
	return errors.New("DTracePrivileges is not yet implemented")
}

func (c *Mod_privileges) dirPrivilegesMode(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: PrivilegesMode directive
	return errors.New("PrivilegesMode is not yet implemented")
}

func (c *Mod_privileges) dirVHostCGIMode(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VHostCGIMode directive
	return errors.New("VHostCGIMode is not yet implemented")
}

func (c *Mod_privileges) dirVHostCGIPrivs(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VHostCGIPrivs directive
	return errors.New("VHostCGIPrivs is not yet implemented")
}

func (c *Mod_privileges) dirVHostGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VHostGroup directive
	return errors.New("VHostGroup is not yet implemented")
}

func (c *Mod_privileges) dirVHostPrivs(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VHostPrivs directive
	return errors.New("VHostPrivs is not yet implemented")
}

func (c *Mod_privileges) dirVHostSecure(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VHostSecure directive
	return errors.New("VHostSecure is not yet implemented")
}

func (c *Mod_privileges) dirVHostUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: VHostUser directive
	return errors.New("VHostUser is not yet implemented")
}
