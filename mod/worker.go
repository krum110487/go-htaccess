// https://httpd.apache.org/docs/2.4/mod/worker.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Worker struct{}

func (Worker) DirCoreDumpDirectory(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CoreDumpDirectory directive
	return errors.New("CoreDumpDirectory is not yet implemented")
}

func (Worker) DirEnableExceptionHook(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: EnableExceptionHook directive
	return errors.New("EnableExceptionHook is not yet implemented")
}

func (Worker) DirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (Worker) DirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (Worker) DirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (Worker) DirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (Worker) DirMaxMemFree(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxMemFree directive
	return errors.New("MaxMemFree is not yet implemented")
}

func (Worker) DirMaxRequestWorkers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxRequestWorkers directive
	return errors.New("MaxRequestWorkers is not yet implemented")
}

func (Worker) DirMaxSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MaxSpareThreads directive
	return errors.New("MaxSpareThreads is not yet implemented")
}

func (Worker) DirMinSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MinSpareThreads directive
	return errors.New("MinSpareThreads is not yet implemented")
}

func (Worker) DirPidFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: PidFile directive
	return errors.New("PidFile is not yet implemented")
}

func (Worker) DirReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ReceiveBufferSize directive
	return errors.New("ReceiveBufferSize is not yet implemented")
}

func (Worker) DirScoreBoardFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ScoreBoardFile directive
	return errors.New("ScoreBoardFile is not yet implemented")
}

func (Worker) DirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (Worker) DirServerLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ServerLimit directive
	return errors.New("ServerLimit is not yet implemented")
}

func (Worker) DirStartServers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: StartServers directive
	return errors.New("StartServers is not yet implemented")
}

func (Worker) DirThreadLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ThreadLimit directive
	return errors.New("ThreadLimit is not yet implemented")
}

func (Worker) DirThreadsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ThreadsPerChild directive
	return errors.New("ThreadsPerChild is not yet implemented")
}

func (Worker) DirThreadStackSize(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ThreadStackSize directive
	return errors.New("ThreadStackSize is not yet implemented")
}

func (Worker) DirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
