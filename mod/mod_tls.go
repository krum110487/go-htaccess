// https://httpd.apache.org/docs/2.4/mod/mod_tls.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_tls struct{}

func (Mod_tls) DirTLSCertificate(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSCertificate directive
	return errors.New("TLSCertificate is not yet implemented")
}

func (Mod_tls) DirTLSCiphersPrefer(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSCiphersPrefer directive
	return errors.New("TLSCiphersPrefer is not yet implemented")
}

func (Mod_tls) DirTLSCiphersSuppress(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSCiphersSuppress directive
	return errors.New("TLSCiphersSuppress is not yet implemented")
}

func (Mod_tls) DirTLSEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSEngine directive
	return errors.New("TLSEngine is not yet implemented")
}

func (Mod_tls) DirTLSHonorClientOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSHonorClientOrder directive
	return errors.New("TLSHonorClientOrder is not yet implemented")
}

func (Mod_tls) DirTLSOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSOptions directive
	return errors.New("TLSOptions is not yet implemented")
}

func (Mod_tls) DirTLSProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSProtocol directive
	return errors.New("TLSProtocol is not yet implemented")
}

func (Mod_tls) DirTLSProxyCA(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSProxyCA directive
	return errors.New("TLSProxyCA is not yet implemented")
}

func (Mod_tls) DirTLSProxyCiphersPrefer(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSProxyCiphersPrefer directive
	return errors.New("TLSProxyCiphersPrefer is not yet implemented")
}

func (Mod_tls) DirTLSProxyCiphersSuppress(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSProxyCiphersSuppress directive
	return errors.New("TLSProxyCiphersSuppress is not yet implemented")
}

func (Mod_tls) DirTLSProxyEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSProxyEngine directive
	return errors.New("TLSProxyEngine is not yet implemented")
}

func (Mod_tls) DirTLSProxyMachineCertificate(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSProxyMachineCertificate directive
	return errors.New("TLSProxyMachineCertificate is not yet implemented")
}

func (Mod_tls) DirTLSProxyProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSProxyProtocol directive
	return errors.New("TLSProxyProtocol is not yet implemented")
}

func (Mod_tls) DirTLSSessionCache(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSSessionCache directive
	return errors.New("TLSSessionCache is not yet implemented")
}

func (Mod_tls) DirTLSStrictSNI(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TLSStrictSNI directive
	return errors.New("TLSStrictSNI is not yet implemented")
}
