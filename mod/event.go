// https://httpd.apache.org/docs/2.4/mod/event.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Event struct{}

func (c *Event) dirAsyncRequestWorkerFactor(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AsyncRequestWorkerFactor directive
	return errors.New("AsyncRequestWorkerFactor is not yet implemented")
}

func (c *Event) dirCoreDumpDirectory(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CoreDumpDirectory directive
	return errors.New("CoreDumpDirectory is not yet implemented")
}

func (c *Event) dirEnableExceptionHook(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: EnableExceptionHook directive
	return errors.New("EnableExceptionHook is not yet implemented")
}

func (c *Event) dirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (c *Event) dirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (c *Event) dirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (c *Event) dirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (c *Event) dirMaxMemFree(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxMemFree directive
	return errors.New("MaxMemFree is not yet implemented")
}

func (c *Event) dirMaxRequestWorkers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRequestWorkers directive
	return errors.New("MaxRequestWorkers is not yet implemented")
}

func (c *Event) dirMaxSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxSpareThreads directive
	return errors.New("MaxSpareThreads is not yet implemented")
}

func (c *Event) dirMinSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MinSpareThreads directive
	return errors.New("MinSpareThreads is not yet implemented")
}

func (c *Event) dirPidFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: PidFile directive
	return errors.New("PidFile is not yet implemented")
}

func (c *Event) dirScoreBoardFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScoreBoardFile directive
	return errors.New("ScoreBoardFile is not yet implemented")
}

func (c *Event) dirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (c *Event) dirServerLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerLimit directive
	return errors.New("ServerLimit is not yet implemented")
}

func (c *Event) dirStartServers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: StartServers directive
	return errors.New("StartServers is not yet implemented")
}

func (c *Event) dirThreadLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ThreadLimit directive
	return errors.New("ThreadLimit is not yet implemented")
}

func (c *Event) dirThreadsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ThreadsPerChild directive
	return errors.New("ThreadsPerChild is not yet implemented")
}

func (c *Event) dirThreadStackSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ThreadStackSize directive
	return errors.New("ThreadStackSize is not yet implemented")
}

func (c *Event) dirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
