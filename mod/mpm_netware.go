// https://httpd.apache.org/docs/2.4/mod/mpm_netware.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mpm_netware struct{}

func (c *Mpm_netware) dirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (c *Mpm_netware) dirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (c *Mpm_netware) dirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (c *Mpm_netware) dirMaxMemFree(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxMemFree directive
	return errors.New("MaxMemFree is not yet implemented")
}

func (c *Mpm_netware) dirMaxSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxSpareThreads directive
	return errors.New("MaxSpareThreads is not yet implemented")
}

func (c *Mpm_netware) dirMaxThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxThreads directive
	return errors.New("MaxThreads is not yet implemented")
}

func (c *Mpm_netware) dirMinSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MinSpareThreads directive
	return errors.New("MinSpareThreads is not yet implemented")
}

func (c *Mpm_netware) dirReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ReceiveBufferSize directive
	return errors.New("ReceiveBufferSize is not yet implemented")
}

func (c *Mpm_netware) dirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (c *Mpm_netware) dirStartThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: StartThreads directive
	return errors.New("StartThreads is not yet implemented")
}

func (c *Mpm_netware) dirThreadStackSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ThreadStackSize directive
	return errors.New("ThreadStackSize is not yet implemented")
}
