// https://httpd.apache.org/docs/2.4/mod/mod_tls.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_tls struct{}

func (c *Mod_tls) dirTLSCertificate(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSCertificate directive
	return errors.New("TLSCertificate is not yet implemented")
}

func (c *Mod_tls) dirTLSCiphersPrefer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSCiphersPrefer directive
	return errors.New("TLSCiphersPrefer is not yet implemented")
}

func (c *Mod_tls) dirTLSCiphersSuppress(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSCiphersSuppress directive
	return errors.New("TLSCiphersSuppress is not yet implemented")
}

func (c *Mod_tls) dirTLSEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSEngine directive
	return errors.New("TLSEngine is not yet implemented")
}

func (c *Mod_tls) dirTLSHonorClientOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSHonorClientOrder directive
	return errors.New("TLSHonorClientOrder is not yet implemented")
}

func (c *Mod_tls) dirTLSOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSOptions directive
	return errors.New("TLSOptions is not yet implemented")
}

func (c *Mod_tls) dirTLSProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSProtocol directive
	return errors.New("TLSProtocol is not yet implemented")
}

func (c *Mod_tls) dirTLSProxyCA(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSProxyCA directive
	return errors.New("TLSProxyCA is not yet implemented")
}

func (c *Mod_tls) dirTLSProxyCiphersPrefer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSProxyCiphersPrefer directive
	return errors.New("TLSProxyCiphersPrefer is not yet implemented")
}

func (c *Mod_tls) dirTLSProxyCiphersSuppress(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSProxyCiphersSuppress directive
	return errors.New("TLSProxyCiphersSuppress is not yet implemented")
}

func (c *Mod_tls) dirTLSProxyEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSProxyEngine directive
	return errors.New("TLSProxyEngine is not yet implemented")
}

func (c *Mod_tls) dirTLSProxyMachineCertificate(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSProxyMachineCertificate directive
	return errors.New("TLSProxyMachineCertificate is not yet implemented")
}

func (c *Mod_tls) dirTLSProxyProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSProxyProtocol directive
	return errors.New("TLSProxyProtocol is not yet implemented")
}

func (c *Mod_tls) dirTLSSessionCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSSessionCache directive
	return errors.New("TLSSessionCache is not yet implemented")
}

func (c *Mod_tls) dirTLSStrictSNI(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TLSStrictSNI directive
	return errors.New("TLSStrictSNI is not yet implemented")
}
