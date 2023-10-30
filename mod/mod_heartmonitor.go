// https://httpd.apache.org/docs/2.4/mod/mod_heartmonitor.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_heartmonitor struct{}

func (c *Mod_heartmonitor) dirHeartbeatListen(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HeartbeatListen directive
	return errors.New("HeartbeatListen is not yet implemented")
}

func (c *Mod_heartmonitor) dirHeartbeatMaxServers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HeartbeatMaxServers directive
	return errors.New("HeartbeatMaxServers is not yet implemented")
}

func (c *Mod_heartmonitor) dirHeartbeatStorage(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HeartbeatStorage directive
	return errors.New("HeartbeatStorage is not yet implemented")
}
