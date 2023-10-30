// https://httpd.apache.org/docs/2.4/mod/mod_ssl.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_ssl struct{}

func (c *Mod_ssl) dirSSLCACertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCACertificateFile directive
	return errors.New("SSLCACertificateFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLCACertificatePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCACertificatePath directive
	return errors.New("SSLCACertificatePath is not yet implemented")
}

func (c *Mod_ssl) dirSSLCADNRequestFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCADNRequestFile directive
	return errors.New("SSLCADNRequestFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLCADNRequestPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCADNRequestPath directive
	return errors.New("SSLCADNRequestPath is not yet implemented")
}

func (c *Mod_ssl) dirSSLCARevocationCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCARevocationCheck directive
	return errors.New("SSLCARevocationCheck is not yet implemented")
}

func (c *Mod_ssl) dirSSLCARevocationFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCARevocationFile directive
	return errors.New("SSLCARevocationFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLCARevocationPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCARevocationPath directive
	return errors.New("SSLCARevocationPath is not yet implemented")
}

func (c *Mod_ssl) dirSSLCertificateChainFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCertificateChainFile directive
	return errors.New("SSLCertificateChainFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCertificateFile directive
	return errors.New("SSLCertificateFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLCertificateKeyFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCertificateKeyFile directive
	return errors.New("SSLCertificateKeyFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLCipherSuite(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCipherSuite directive
	return errors.New("SSLCipherSuite is not yet implemented")
}

func (c *Mod_ssl) dirSSLCompression(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCompression directive
	return errors.New("SSLCompression is not yet implemented")
}

func (c *Mod_ssl) dirSSLCryptoDevice(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLCryptoDevice directive
	return errors.New("SSLCryptoDevice is not yet implemented")
}

func (c *Mod_ssl) dirSSLEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLEngine directive
	return errors.New("SSLEngine is not yet implemented")
}

func (c *Mod_ssl) dirSSLFIPS(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLFIPS directive
	return errors.New("SSLFIPS is not yet implemented")
}

func (c *Mod_ssl) dirSSLHonorCipherOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLHonorCipherOrder directive
	return errors.New("SSLHonorCipherOrder is not yet implemented")
}

func (c *Mod_ssl) dirSSLInsecureRenegotiation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLInsecureRenegotiation directive
	return errors.New("SSLInsecureRenegotiation is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPDefaultResponder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPDefaultResponder directive
	return errors.New("SSLOCSPDefaultResponder is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPEnable directive
	return errors.New("SSLOCSPEnable is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPNoverify(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPNoverify directive
	return errors.New("SSLOCSPNoverify is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPOverrideResponder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPOverrideResponder directive
	return errors.New("SSLOCSPOverrideResponder is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPProxyURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPProxyURL directive
	return errors.New("SSLOCSPProxyURL is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPResponderCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponderCertificateFile directive
	return errors.New("SSLOCSPResponderCertificateFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPResponderTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponderTimeout directive
	return errors.New("SSLOCSPResponderTimeout is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPResponseMaxAge(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponseMaxAge directive
	return errors.New("SSLOCSPResponseMaxAge is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPResponseTimeSkew(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPResponseTimeSkew directive
	return errors.New("SSLOCSPResponseTimeSkew is not yet implemented")
}

func (c *Mod_ssl) dirSSLOCSPUseRequestNonce(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOCSPUseRequestNonce directive
	return errors.New("SSLOCSPUseRequestNonce is not yet implemented")
}

func (c *Mod_ssl) dirSSLOpenSSLConfCmd(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOpenSSLConfCmd directive
	return errors.New("SSLOpenSSLConfCmd is not yet implemented")
}

func (c *Mod_ssl) dirSSLOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLOptions directive
	return errors.New("SSLOptions is not yet implemented")
}

func (c *Mod_ssl) dirSSLPassPhraseDialog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLPassPhraseDialog directive
	return errors.New("SSLPassPhraseDialog is not yet implemented")
}

func (c *Mod_ssl) dirSSLProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProtocol directive
	return errors.New("SSLProtocol is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCACertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCACertificateFile directive
	return errors.New("SSLProxyCACertificateFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCACertificatePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCACertificatePath directive
	return errors.New("SSLProxyCACertificatePath is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCARevocationCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCARevocationCheck directive
	return errors.New("SSLProxyCARevocationCheck is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCARevocationFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCARevocationFile directive
	return errors.New("SSLProxyCARevocationFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCARevocationPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCARevocationPath directive
	return errors.New("SSLProxyCARevocationPath is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCheckPeerCN(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCheckPeerCN directive
	return errors.New("SSLProxyCheckPeerCN is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCheckPeerExpire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCheckPeerExpire directive
	return errors.New("SSLProxyCheckPeerExpire is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCheckPeerName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCheckPeerName directive
	return errors.New("SSLProxyCheckPeerName is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyCipherSuite(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyCipherSuite directive
	return errors.New("SSLProxyCipherSuite is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyEngine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyEngine directive
	return errors.New("SSLProxyEngine is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyMachineCertificateChainFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyMachineCertificateChainFile directive
	return errors.New("SSLProxyMachineCertificateChainFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyMachineCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyMachineCertificateFile directive
	return errors.New("SSLProxyMachineCertificateFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyMachineCertificatePath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyMachineCertificatePath directive
	return errors.New("SSLProxyMachineCertificatePath is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyProtocol directive
	return errors.New("SSLProxyProtocol is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyVerify(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyVerify directive
	return errors.New("SSLProxyVerify is not yet implemented")
}

func (c *Mod_ssl) dirSSLProxyVerifyDepth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLProxyVerifyDepth directive
	return errors.New("SSLProxyVerifyDepth is not yet implemented")
}

func (c *Mod_ssl) dirSSLRandomSeed(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRandomSeed directive
	return errors.New("SSLRandomSeed is not yet implemented")
}

func (c *Mod_ssl) dirSSLRenegBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRenegBufferSize directive
	return errors.New("SSLRenegBufferSize is not yet implemented")
}

func (c *Mod_ssl) dirSSLRequire(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRequire directive
	return errors.New("SSLRequire is not yet implemented")
}

func (c *Mod_ssl) dirSSLRequireSSL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLRequireSSL directive
	return errors.New("SSLRequireSSL is not yet implemented")
}

func (c *Mod_ssl) dirSSLSessionCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionCache directive
	return errors.New("SSLSessionCache is not yet implemented")
}

func (c *Mod_ssl) dirSSLSessionCacheTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionCacheTimeout directive
	return errors.New("SSLSessionCacheTimeout is not yet implemented")
}

func (c *Mod_ssl) dirSSLSessionTicketKeyFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionTicketKeyFile directive
	return errors.New("SSLSessionTicketKeyFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLSessionTickets(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSessionTickets directive
	return errors.New("SSLSessionTickets is not yet implemented")
}

func (c *Mod_ssl) dirSSLSRPUnknownUserSeed(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSRPUnknownUserSeed directive
	return errors.New("SSLSRPUnknownUserSeed is not yet implemented")
}

func (c *Mod_ssl) dirSSLSRPVerifierFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLSRPVerifierFile directive
	return errors.New("SSLSRPVerifierFile is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingCache(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingCache directive
	return errors.New("SSLStaplingCache is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingErrorCacheTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingErrorCacheTimeout directive
	return errors.New("SSLStaplingErrorCacheTimeout is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingFakeTryLater(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingFakeTryLater directive
	return errors.New("SSLStaplingFakeTryLater is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingForceURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingForceURL directive
	return errors.New("SSLStaplingForceURL is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingResponderTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingResponderTimeout directive
	return errors.New("SSLStaplingResponderTimeout is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingResponseMaxAge(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingResponseMaxAge directive
	return errors.New("SSLStaplingResponseMaxAge is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingResponseTimeSkew(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingResponseTimeSkew directive
	return errors.New("SSLStaplingResponseTimeSkew is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingReturnResponderErrors(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingReturnResponderErrors directive
	return errors.New("SSLStaplingReturnResponderErrors is not yet implemented")
}

func (c *Mod_ssl) dirSSLStaplingStandardCacheTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStaplingStandardCacheTimeout directive
	return errors.New("SSLStaplingStandardCacheTimeout is not yet implemented")
}

func (c *Mod_ssl) dirSSLStrictSNIVHostCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLStrictSNIVHostCheck directive
	return errors.New("SSLStrictSNIVHostCheck is not yet implemented")
}

func (c *Mod_ssl) dirSSLUserName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLUserName directive
	return errors.New("SSLUserName is not yet implemented")
}

func (c *Mod_ssl) dirSSLUseStapling(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLUseStapling directive
	return errors.New("SSLUseStapling is not yet implemented")
}

func (c *Mod_ssl) dirSSLVerifyClient(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLVerifyClient directive
	return errors.New("SSLVerifyClient is not yet implemented")
}

func (c *Mod_ssl) dirSSLVerifyDepth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SSLVerifyDepth directive
	return errors.New("SSLVerifyDepth is not yet implemented")
}
