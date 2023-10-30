// https://httpd.apache.org/docs/2.4/mod/mod_ssl.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_ssl struct{}

func (c *Mod_ssl) DirSSLCACertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCACertificateFile directive
	return errors.New("SSLCACertificateFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLCACertificatePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCACertificatePath directive
	return errors.New("SSLCACertificatePath is not yet implemented")
}

func (c *Mod_ssl) DirSSLCADNRequestFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCADNRequestFile directive
	return errors.New("SSLCADNRequestFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLCADNRequestPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCADNRequestPath directive
	return errors.New("SSLCADNRequestPath is not yet implemented")
}

func (c *Mod_ssl) DirSSLCARevocationCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCARevocationCheck directive
	return errors.New("SSLCARevocationCheck is not yet implemented")
}

func (c *Mod_ssl) DirSSLCARevocationFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCARevocationFile directive
	return errors.New("SSLCARevocationFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLCARevocationPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCARevocationPath directive
	return errors.New("SSLCARevocationPath is not yet implemented")
}

func (c *Mod_ssl) DirSSLCertificateChainFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCertificateChainFile directive
	return errors.New("SSLCertificateChainFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCertificateFile directive
	return errors.New("SSLCertificateFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLCertificateKeyFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCertificateKeyFile directive
	return errors.New("SSLCertificateKeyFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLCipherSuite(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCipherSuite directive
	return errors.New("SSLCipherSuite is not yet implemented")
}

func (c *Mod_ssl) DirSSLCompression(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCompression directive
	return errors.New("SSLCompression is not yet implemented")
}

func (c *Mod_ssl) DirSSLCryptoDevice(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCryptoDevice directive
	return errors.New("SSLCryptoDevice is not yet implemented")
}

func (c *Mod_ssl) DirSSLEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLEngine directive
	return errors.New("SSLEngine is not yet implemented")
}

func (c *Mod_ssl) DirSSLFIPS(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLFIPS directive
	return errors.New("SSLFIPS is not yet implemented")
}

func (c *Mod_ssl) DirSSLHonorCipherOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLHonorCipherOrder directive
	return errors.New("SSLHonorCipherOrder is not yet implemented")
}

func (c *Mod_ssl) DirSSLInsecureRenegotiation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLInsecureRenegotiation directive
	return errors.New("SSLInsecureRenegotiation is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPDefaultResponder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPDefaultResponder directive
	return errors.New("SSLOCSPDefaultResponder is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPEnable directive
	return errors.New("SSLOCSPEnable is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPNoverify(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPNoverify directive
	return errors.New("SSLOCSPNoverify is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPOverrideResponder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPOverrideResponder directive
	return errors.New("SSLOCSPOverrideResponder is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPProxyURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPProxyURL directive
	return errors.New("SSLOCSPProxyURL is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPResponderCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponderCertificateFile directive
	return errors.New("SSLOCSPResponderCertificateFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPResponderTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponderTimeout directive
	return errors.New("SSLOCSPResponderTimeout is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPResponseMaxAge(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponseMaxAge directive
	return errors.New("SSLOCSPResponseMaxAge is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPResponseTimeSkew(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponseTimeSkew directive
	return errors.New("SSLOCSPResponseTimeSkew is not yet implemented")
}

func (c *Mod_ssl) DirSSLOCSPUseRequestNonce(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPUseRequestNonce directive
	return errors.New("SSLOCSPUseRequestNonce is not yet implemented")
}

func (c *Mod_ssl) DirSSLOpenSSLConfCmd(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOpenSSLConfCmd directive
	return errors.New("SSLOpenSSLConfCmd is not yet implemented")
}

func (c *Mod_ssl) DirSSLOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOptions directive
	return errors.New("SSLOptions is not yet implemented")
}

func (c *Mod_ssl) DirSSLPassPhraseDialog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLPassPhraseDialog directive
	return errors.New("SSLPassPhraseDialog is not yet implemented")
}

func (c *Mod_ssl) DirSSLProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProtocol directive
	return errors.New("SSLProtocol is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCACertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCACertificateFile directive
	return errors.New("SSLProxyCACertificateFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCACertificatePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCACertificatePath directive
	return errors.New("SSLProxyCACertificatePath is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCARevocationCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCARevocationCheck directive
	return errors.New("SSLProxyCARevocationCheck is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCARevocationFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCARevocationFile directive
	return errors.New("SSLProxyCARevocationFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCARevocationPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCARevocationPath directive
	return errors.New("SSLProxyCARevocationPath is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCheckPeerCN(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCheckPeerCN directive
	return errors.New("SSLProxyCheckPeerCN is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCheckPeerExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCheckPeerExpire directive
	return errors.New("SSLProxyCheckPeerExpire is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCheckPeerName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCheckPeerName directive
	return errors.New("SSLProxyCheckPeerName is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyCipherSuite(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCipherSuite directive
	return errors.New("SSLProxyCipherSuite is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyEngine directive
	return errors.New("SSLProxyEngine is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyMachineCertificateChainFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyMachineCertificateChainFile directive
	return errors.New("SSLProxyMachineCertificateChainFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyMachineCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyMachineCertificateFile directive
	return errors.New("SSLProxyMachineCertificateFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyMachineCertificatePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyMachineCertificatePath directive
	return errors.New("SSLProxyMachineCertificatePath is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyProtocol directive
	return errors.New("SSLProxyProtocol is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyVerify(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyVerify directive
	return errors.New("SSLProxyVerify is not yet implemented")
}

func (c *Mod_ssl) DirSSLProxyVerifyDepth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyVerifyDepth directive
	return errors.New("SSLProxyVerifyDepth is not yet implemented")
}

func (c *Mod_ssl) DirSSLRandomSeed(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRandomSeed directive
	return errors.New("SSLRandomSeed is not yet implemented")
}

func (c *Mod_ssl) DirSSLRenegBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRenegBufferSize directive
	return errors.New("SSLRenegBufferSize is not yet implemented")
}

func (c *Mod_ssl) DirSSLRequire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRequire directive
	return errors.New("SSLRequire is not yet implemented")
}

func (c *Mod_ssl) DirSSLRequireSSL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRequireSSL directive
	return errors.New("SSLRequireSSL is not yet implemented")
}

func (c *Mod_ssl) DirSSLSessionCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionCache directive
	return errors.New("SSLSessionCache is not yet implemented")
}

func (c *Mod_ssl) DirSSLSessionCacheTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionCacheTimeout directive
	return errors.New("SSLSessionCacheTimeout is not yet implemented")
}

func (c *Mod_ssl) DirSSLSessionTicketKeyFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionTicketKeyFile directive
	return errors.New("SSLSessionTicketKeyFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLSessionTickets(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionTickets directive
	return errors.New("SSLSessionTickets is not yet implemented")
}

func (c *Mod_ssl) DirSSLSRPUnknownUserSeed(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSRPUnknownUserSeed directive
	return errors.New("SSLSRPUnknownUserSeed is not yet implemented")
}

func (c *Mod_ssl) DirSSLSRPVerifierFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSRPVerifierFile directive
	return errors.New("SSLSRPVerifierFile is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingCache directive
	return errors.New("SSLStaplingCache is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingErrorCacheTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingErrorCacheTimeout directive
	return errors.New("SSLStaplingErrorCacheTimeout is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingFakeTryLater(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingFakeTryLater directive
	return errors.New("SSLStaplingFakeTryLater is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingForceURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingForceURL directive
	return errors.New("SSLStaplingForceURL is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingResponderTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingResponderTimeout directive
	return errors.New("SSLStaplingResponderTimeout is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingResponseMaxAge(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingResponseMaxAge directive
	return errors.New("SSLStaplingResponseMaxAge is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingResponseTimeSkew(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingResponseTimeSkew directive
	return errors.New("SSLStaplingResponseTimeSkew is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingReturnResponderErrors(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingReturnResponderErrors directive
	return errors.New("SSLStaplingReturnResponderErrors is not yet implemented")
}

func (c *Mod_ssl) DirSSLStaplingStandardCacheTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingStandardCacheTimeout directive
	return errors.New("SSLStaplingStandardCacheTimeout is not yet implemented")
}

func (c *Mod_ssl) DirSSLStrictSNIVHostCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStrictSNIVHostCheck directive
	return errors.New("SSLStrictSNIVHostCheck is not yet implemented")
}

func (c *Mod_ssl) DirSSLUserName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLUserName directive
	return errors.New("SSLUserName is not yet implemented")
}

func (c *Mod_ssl) DirSSLUseStapling(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLUseStapling directive
	return errors.New("SSLUseStapling is not yet implemented")
}

func (c *Mod_ssl) DirSSLVerifyClient(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLVerifyClient directive
	return errors.New("SSLVerifyClient is not yet implemented")
}

func (c *Mod_ssl) DirSSLVerifyDepth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLVerifyDepth directive
	return errors.New("SSLVerifyDepth is not yet implemented")
}
