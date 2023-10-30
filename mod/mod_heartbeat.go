// https://httpd.apache.org/docs/2.4/mod/mod_heartbeat.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_heartbeat struct{}

func (c *Mod_heartbeat) dirHeartbeatAddress(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HeartbeatAddress directive
	return errors.New("HeartbeatAddress is not yet implemented")
}
