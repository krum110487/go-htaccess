// https://httpd.apache.org/docs/2.4/mod/core.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Prefork struct{}

func (c *Prefork) dirCoreDumpDirectory(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CoreDumpDirectory directive
	return errors.New("CoreDumpDirectory is not yet implemented")
}

func (c *Prefork) dirEnableExceptionHook(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: EnableExceptionHook directive
	return errors.New("EnableExceptionHook is not yet implemented")
}

func (c *Prefork) dirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (c *Prefork) dirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (c *Prefork) dirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (c *Prefork) dirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (c *Prefork) dirMaxMemFree(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxMemFree directive
	return errors.New("MaxMemFree is not yet implemented")
}

func (c *Prefork) dirMaxRequestWorkers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRequestWorkers directive
	return errors.New("MaxRequestWorkers is not yet implemented")
}

func (c *Prefork) dirMaxSpareServers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxSpareServers directive
	return errors.New("MaxSpareServers is not yet implemented")
}

func (c *Prefork) dirMinSpareServers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MinSpareServers directive
	return errors.New("MinSpareServers is not yet implemented")
}

func (c *Prefork) dirPidFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: PidFile directive
	return errors.New("PidFile is not yet implemented")
}

func (c *Prefork) dirReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ReceiveBufferSize directive
	return errors.New("ReceiveBufferSize is not yet implemented")
}

func (c *Prefork) dirScoreBoardFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScoreBoardFile directive
	return errors.New("ScoreBoardFile is not yet implemented")
}

func (c *Prefork) dirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (c *Prefork) dirServerLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerLimit directive
	return errors.New("ServerLimit is not yet implemented")
}

func (c *Prefork) dirStartServers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: StartServers directive
	return errors.New("StartServers is not yet implemented")
}

func (c *Prefork) dirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
