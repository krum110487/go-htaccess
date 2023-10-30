// https://httpd.apache.org/docs/2.4/mod/mod_watchdog.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_watchdog struct{}

func (c *Mod_watchdog) dirWatchdogInterval(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: WatchdogInterval directive
	return errors.New("WatchdogInterval is not yet implemented")
}
