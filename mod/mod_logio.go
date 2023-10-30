// https://httpd.apache.org/docs/2.4/mod/mod_logio.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_logio struct{}

func (c *Mod_logio) DirLogIOTrackTTFB(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LogIOTrackTTFB directive
	return errors.New("LogIOTrackTTFB is not yet implemented")
}
