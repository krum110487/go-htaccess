// https://httpd.apache.org/docs/2.4/mod/mod_heartmonitor.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_heartmonitor struct{}

func (Mod_heartmonitor) DirHeartbeatListen(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: HeartbeatListen directive
	return errors.New("HeartbeatListen is not yet implemented")
}

func (Mod_heartmonitor) DirHeartbeatMaxServers(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: HeartbeatMaxServers directive
	return errors.New("HeartbeatMaxServers is not yet implemented")
}

func (Mod_heartmonitor) DirHeartbeatStorage(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: HeartbeatStorage directive
	return errors.New("HeartbeatStorage is not yet implemented")
}
