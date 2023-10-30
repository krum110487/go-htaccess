// https://httpd.apache.org/docs/2.4/mod/mod_lbmethod_heartbeat.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_lbmethod_heartbeat struct{}

func (c *Mod_lbmethod_heartbeat) DirHeartbeatStorage(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HeartbeatStorage directive
	return errors.New("HeartbeatStorage is not yet implemented")
}
