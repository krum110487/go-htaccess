// https://httpd.apache.org/docs/2.4/mod/mpm_netware.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mpm_netware struct{}

func (Mpm_netware) DirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (Mpm_netware) DirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (Mpm_netware) DirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (Mpm_netware) DirMaxMemFree(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxMemFree directive
	return errors.New("MaxMemFree is not yet implemented")
}

func (Mpm_netware) DirMaxSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxSpareThreads directive
	return errors.New("MaxSpareThreads is not yet implemented")
}

func (Mpm_netware) DirMaxThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxThreads directive
	return errors.New("MaxThreads is not yet implemented")
}

func (Mpm_netware) DirMinSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MinSpareThreads directive
	return errors.New("MinSpareThreads is not yet implemented")
}

func (Mpm_netware) DirReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ReceiveBufferSize directive
	return errors.New("ReceiveBufferSize is not yet implemented")
}

func (Mpm_netware) DirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (Mpm_netware) DirStartThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: StartThreads directive
	return errors.New("StartThreads is not yet implemented")
}

func (Mpm_netware) DirThreadStackSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ThreadStackSize directive
	return errors.New("ThreadStackSize is not yet implemented")
}
