// https://httpd.apache.org/docs/2.4/mod/event.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Event struct{}

func (Event) DirAsyncRequestWorkerFactor(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AsyncRequestWorkerFactor directive
	return errors.New("AsyncRequestWorkerFactor is not yet implemented")
}

func (Event) DirCoreDumpDirectory(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CoreDumpDirectory directive
	return errors.New("CoreDumpDirectory is not yet implemented")
}

func (Event) DirEnableExceptionHook(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: EnableExceptionHook directive
	return errors.New("EnableExceptionHook is not yet implemented")
}

func (Event) DirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (Event) DirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (Event) DirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (Event) DirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (Event) DirMaxMemFree(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxMemFree directive
	return errors.New("MaxMemFree is not yet implemented")
}

func (Event) DirMaxRequestWorkers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxRequestWorkers directive
	return errors.New("MaxRequestWorkers is not yet implemented")
}

func (Event) DirMaxSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxSpareThreads directive
	return errors.New("MaxSpareThreads is not yet implemented")
}

func (Event) DirMinSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MinSpareThreads directive
	return errors.New("MinSpareThreads is not yet implemented")
}

func (Event) DirPidFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: PidFile directive
	return errors.New("PidFile is not yet implemented")
}

func (Event) DirScoreBoardFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ScoreBoardFile directive
	return errors.New("ScoreBoardFile is not yet implemented")
}

func (Event) DirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (Event) DirServerLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ServerLimit directive
	return errors.New("ServerLimit is not yet implemented")
}

func (Event) DirStartServers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: StartServers directive
	return errors.New("StartServers is not yet implemented")
}

func (Event) DirThreadLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ThreadLimit directive
	return errors.New("ThreadLimit is not yet implemented")
}

func (Event) DirThreadsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ThreadsPerChild directive
	return errors.New("ThreadsPerChild is not yet implemented")
}

func (Event) DirThreadStackSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ThreadStackSize directive
	return errors.New("ThreadStackSize is not yet implemented")
}

func (Event) DirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
