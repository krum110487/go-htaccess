// https://httpd.apache.org/docs/2.4/mod/mod_log_config.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_log_config struct{}

func (Mod_log_config) DirBufferedLogs(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: BufferedLogs directive
	return errors.New("BufferedLogs is not yet implemented")
}

func (Mod_log_config) DirCustomLog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: CustomLog directive
	return errors.New("CustomLog is not yet implemented")
}

func (Mod_log_config) DirGlobalLog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: GlobalLog directive
	return errors.New("GlobalLog is not yet implemented")
}

func (Mod_log_config) DirLogFormat(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: LogFormat directive
	return errors.New("LogFormat is not yet implemented")
}

func (Mod_log_config) DirTransferLog(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TransferLog directive
	return errors.New("TransferLog is not yet implemented")
}
