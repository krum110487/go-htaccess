// https://httpd.apache.org/docs/2.4/mod/mpmt_os2.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mpmt_os2 struct{}

func (Mpmt_os2) DirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (Mpmt_os2) DirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (Mpmt_os2) DirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (Mpmt_os2) DirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (Mpmt_os2) DirMaxSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxSpareThreads directive
	return errors.New("MaxSpareThreads is not yet implemented")
}

func (Mpmt_os2) DirMinSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MinSpareThreads directive
	return errors.New("MinSpareThreads is not yet implemented")
}

func (Mpmt_os2) DirPidFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: PidFile directive
	return errors.New("PidFile is not yet implemented")
}

func (Mpmt_os2) DirReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ReceiveBufferSize directive
	return errors.New("ReceiveBufferSize is not yet implemented")
}

func (Mpmt_os2) DirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (Mpmt_os2) DirStartServers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: StartServers directive
	return errors.New("StartServers is not yet implemented")
}

func (Mpmt_os2) DirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
