// https://httpd.apache.org/docs/2.4/mod/mod_http2.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_http2 struct{}

func (c *Mod_http2) DirH2CopyFiles(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2CopyFiles directive
	return errors.New("H2CopyFiles is not yet implemented")
}

func (c *Mod_http2) DirH2Direct(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2Direct directive
	return errors.New("H2Direct is not yet implemented")
}

func (c *Mod_http2) DirH2EarlyHint(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2EarlyHint directive
	return errors.New("H2EarlyHint is not yet implemented")
}

func (c *Mod_http2) DirH2EarlyHints(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2EarlyHints directive
	return errors.New("H2EarlyHints is not yet implemented")
}

func (c *Mod_http2) DirH2MaxDataFrameLen(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2MaxDataFrameLen directive
	return errors.New("H2MaxDataFrameLen is not yet implemented")
}

func (c *Mod_http2) DirH2MaxSessionStreams(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2MaxSessionStreams directive
	return errors.New("H2MaxSessionStreams is not yet implemented")
}

func (c *Mod_http2) DirH2MaxWorkerIdleSeconds(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2MaxWorkerIdleSeconds directive
	return errors.New("H2MaxWorkerIdleSeconds is not yet implemented")
}

func (c *Mod_http2) DirH2MaxWorkers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2MaxWorkers directive
	return errors.New("H2MaxWorkers is not yet implemented")
}

func (c *Mod_http2) DirH2MinWorkers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2MinWorkers directive
	return errors.New("H2MinWorkers is not yet implemented")
}

func (c *Mod_http2) DirH2ModernTLSOnly(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2ModernTLSOnly directive
	return errors.New("H2ModernTLSOnly is not yet implemented")
}

func (c *Mod_http2) DirH2OutputBuffering(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2OutputBuffering directive
	return errors.New("H2OutputBuffering is not yet implemented")
}

func (c *Mod_http2) DirH2Padding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2Padding directive
	return errors.New("H2Padding is not yet implemented")
}

func (c *Mod_http2) DirH2ProxyRequests(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2ProxyRequests directive
	return errors.New("H2ProxyRequests is not yet implemented")
}

func (c *Mod_http2) DirH2Push(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2Push directive
	return errors.New("H2Push is not yet implemented")
}

func (c *Mod_http2) DirH2PushDiarySize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2PushDiarySize directive
	return errors.New("H2PushDiarySize is not yet implemented")
}

func (c *Mod_http2) DirH2PushPriority(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2PushPriority directive
	return errors.New("H2PushPriority is not yet implemented")
}

func (c *Mod_http2) DirH2PushResource(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2PushResource directive
	return errors.New("H2PushResource is not yet implemented")
}

func (c *Mod_http2) DirH2SerializeHeaders(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2SerializeHeaders directive
	return errors.New("H2SerializeHeaders is not yet implemented")
}

func (c *Mod_http2) DirH2StreamMaxMemSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2StreamMaxMemSize directive
	return errors.New("H2StreamMaxMemSize is not yet implemented")
}

func (c *Mod_http2) DirH2StreamTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2StreamTimeout directive
	return errors.New("H2StreamTimeout is not yet implemented")
}

func (c *Mod_http2) DirH2TLSCoolDownSecs(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2TLSCoolDownSecs directive
	return errors.New("H2TLSCoolDownSecs is not yet implemented")
}

func (c *Mod_http2) DirH2TLSWarmUpSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2TLSWarmUpSize directive
	return errors.New("H2TLSWarmUpSize is not yet implemented")
}

func (c *Mod_http2) DirH2Upgrade(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2Upgrade directive
	return errors.New("H2Upgrade is not yet implemented")
}

func (c *Mod_http2) DirH2WebSockets(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2WebSockets directive
	return errors.New("H2WebSockets is not yet implemented")
}

func (c *Mod_http2) DirH2WindowSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: H2WindowSize directive
	return errors.New("H2WindowSize is not yet implemented")
}
