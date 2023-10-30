// https://httpd.apache.org/docs/2.4/mod/mod_log_config.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_log_config struct{}

func (c *Mod_log_config) dirBufferedLogs(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: BufferedLogs directive
	return errors.New("BufferedLogs is not yet implemented")
}

func (c *Mod_log_config) dirCustomLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CustomLog directive
	return errors.New("CustomLog is not yet implemented")
}

func (c *Mod_log_config) dirGlobalLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: GlobalLog directive
	return errors.New("GlobalLog is not yet implemented")
}

func (c *Mod_log_config) dirLogFormat(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LogFormat directive
	return errors.New("LogFormat is not yet implemented")
}

func (c *Mod_log_config) dirTransferLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TransferLog directive
	return errors.New("TransferLog is not yet implemented")
}
