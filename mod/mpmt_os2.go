// https://httpd.apache.org/docs/2.4/mod/mpmt_os2.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mpmt_os2 struct{}

func (c *Mpmt_os2) dirGroup(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Group directive
	return errors.New("Group is not yet implemented")
}

func (c *Mpmt_os2) dirListen(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Listen directive
	return errors.New("Listen is not yet implemented")
}

func (c *Mpmt_os2) dirListenBacklog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ListenBacklog directive
	return errors.New("ListenBacklog is not yet implemented")
}

func (c *Mpmt_os2) dirMaxConnectionsPerChild(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxConnectionsPerChild directive
	return errors.New("MaxConnectionsPerChild is not yet implemented")
}

func (c *Mpmt_os2) dirMaxSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxSpareThreads directive
	return errors.New("MaxSpareThreads is not yet implemented")
}

func (c *Mpmt_os2) dirMinSpareThreads(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MinSpareThreads directive
	return errors.New("MinSpareThreads is not yet implemented")
}

func (c *Mpmt_os2) dirPidFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: PidFile directive
	return errors.New("PidFile is not yet implemented")
}

func (c *Mpmt_os2) dirReceiveBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ReceiveBufferSize directive
	return errors.New("ReceiveBufferSize is not yet implemented")
}

func (c *Mpmt_os2) dirSendBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SendBufferSize directive
	return errors.New("SendBufferSize is not yet implemented")
}

func (c *Mpmt_os2) dirStartServers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: StartServers directive
	return errors.New("StartServers is not yet implemented")
}

func (c *Mpmt_os2) dirUser(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: User directive
	return errors.New("User is not yet implemented")
}
