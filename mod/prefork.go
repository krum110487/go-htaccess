// https://httpd.apache.org/docs/2.4/mod/core.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Prefork struct{}

func (Prefork) DirCoreDumpDirectory(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CoreDumpDirectory directive
	return errors.New("CoreDumpDirectory is not yet implemented")
}

func (Prefork) DirEnableExceptionHook(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: EnableExceptionHook directive
	return errors.New("EnableExceptionHook is not yet implemented")
}

func (Prefork) DirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (Prefork) DirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (Prefork) DirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (Prefork) DirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (Prefork) DirMaxMemFree(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxMemFree directive
	return errors.New("MaxMemFree is not yet implemented")
}

func (Prefork) DirMaxRequestWorkers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxRequestWorkers directive
	return errors.New("MaxRequestWorkers is not yet implemented")
}

func (Prefork) DirMaxSpareServers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxSpareServers directive
	return errors.New("MaxSpareServers is not yet implemented")
}

func (Prefork) DirMinSpareServers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MinSpareServers directive
	return errors.New("MinSpareServers is not yet implemented")
}

func (Prefork) DirPidFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: PidFile directive
	return errors.New("PidFile is not yet implemented")
}

func (Prefork) DirReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ReceiveBufferSize directive
	return errors.New("ReceiveBufferSize is not yet implemented")
}

func (Prefork) DirScoreBoardFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ScoreBoardFile directive
	return errors.New("ScoreBoardFile is not yet implemented")
}

func (Prefork) DirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (Prefork) DirServerLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ServerLimit directive
	return errors.New("ServerLimit is not yet implemented")
}

func (Prefork) DirStartServers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: StartServers directive
	return errors.New("StartServers is not yet implemented")
}

func (Prefork) DirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
