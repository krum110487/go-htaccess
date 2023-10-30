// https://httpd.apache.org/docs/2.4/mod/mod_md.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_md struct{}

func (c *Mod_md) dirMDActivationDelay(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDActivationDelay directive
	return errors.New("MDActivationDelay is not yet implemented")
}

func (c *Mod_md) dirMDBaseServer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDBaseServer directive
	return errors.New("MDBaseServer is not yet implemented")
}

func (c *Mod_md) dirMDCAChallenges(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCAChallenges directive
	return errors.New("MDCAChallenges is not yet implemented")
}

func (c *Mod_md) dirMDCertificateAgreement(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateAgreement directive
	return errors.New("MDCertificateAgreement is not yet implemented")
}

func (c *Mod_md) dirMDCertificateAuthority(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateAuthority directive
	return errors.New("MDCertificateAuthority is not yet implemented")
}

func (c *Mod_md) dirMDCertificateCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateCheck directive
	return errors.New("MDCertificateCheck is not yet implemented")
}

func (c *Mod_md) dirMDCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateFile directive
	return errors.New("MDCertificateFile is not yet implemented")
}

func (c *Mod_md) dirMDCertificateKeyFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateKeyFile directive
	return errors.New("MDCertificateKeyFile is not yet implemented")
}

func (c *Mod_md) dirMDCertificateMonitor(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateMonitor directive
	return errors.New("MDCertificateMonitor is not yet implemented")
}

func (c *Mod_md) dirMDCertificateProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateProtocol directive
	return errors.New("MDCertificateProtocol is not yet implemented")
}

func (c *Mod_md) dirMDCertificateStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateStatus directive
	return errors.New("MDCertificateStatus is not yet implemented")
}

func (c *Mod_md) dirMDChallengeDns01(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDChallengeDns01 directive
	return errors.New("MDChallengeDns01 is not yet implemented")
}

func (c *Mod_md) dirMDChallengeDns01Version(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDChallengeDns01Version directive
	return errors.New("MDChallengeDns01Version is not yet implemented")
}

func (c *Mod_md) dirMDContactEmail(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDContactEmail directive
	return errors.New("MDContactEmail is not yet implemented")
}

func (c *Mod_md) dirMDDriveMode(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDDriveMode directive
	return errors.New("MDDriveMode is not yet implemented")
}

func (c *Mod_md) dirMDExternalAccountBinding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDExternalAccountBinding directive
	return errors.New("MDExternalAccountBinding is not yet implemented")
}

func (c *Mod_md) dirMDHttpProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDHttpProxy directive
	return errors.New("MDHttpProxy is not yet implemented")
}

func (c *Mod_md) dirMDMatchNames(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMatchNames directive
	return errors.New("MDMatchNames is not yet implemented")
}

func (c *Mod_md) dirMDMember(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMember directive
	return errors.New("MDMember is not yet implemented")
}

func (c *Mod_md) dirMDMembers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMembers directive
	return errors.New("MDMembers is not yet implemented")
}

func (c *Mod_md) dirMDMessageCmd(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMessageCmd directive
	return errors.New("MDMessageCmd is not yet implemented")
}

func (c *Mod_md) dirMDMustStaple(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMustStaple directive
	return errors.New("MDMustStaple is not yet implemented")
}

func (c *Mod_md) dirMDNotifyCmd(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDNotifyCmd directive
	return errors.New("MDNotifyCmd is not yet implemented")
}

func (c *Mod_md) dirMDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDomain directive
	return errors.New("MDomain is not yet implemented")
}

func (c *Mod_md) dirMDomainSet(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <MDomainSet> directive
	return errors.New("<MDomainSet> is not yet implemented")
}

func (c *Mod_md) dirMDPortMap(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDPortMap directive
	return errors.New("MDPortMap is not yet implemented")
}

func (c *Mod_md) dirMDPrivateKeys(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDPrivateKeys directive
	return errors.New("MDPrivateKeys is not yet implemented")
}

func (c *Mod_md) dirMDRenewMode(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRenewMode directive
	return errors.New("MDRenewMode is not yet implemented")
}

func (c *Mod_md) dirMDRenewWindow(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRenewWindow directive
	return errors.New("MDRenewWindow is not yet implemented")
}

func (c *Mod_md) dirMDRequireHttps(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRequireHttps directive
	return errors.New("MDRequireHttps is not yet implemented")
}

func (c *Mod_md) dirMDRetryDelay(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRetryDelay directive
	return errors.New("MDRetryDelay is not yet implemented")
}

func (c *Mod_md) dirMDRetryFailover(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRetryFailover directive
	return errors.New("MDRetryFailover is not yet implemented")
}

func (c *Mod_md) dirMDServerStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDServerStatus directive
	return errors.New("MDServerStatus is not yet implemented")
}

func (c *Mod_md) dirMDStapleOthers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStapleOthers directive
	return errors.New("MDStapleOthers is not yet implemented")
}

func (c *Mod_md) dirMDStapling(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStapling directive
	return errors.New("MDStapling is not yet implemented")
}

func (c *Mod_md) dirMDStaplingKeepResponse(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStaplingKeepResponse directive
	return errors.New("MDStaplingKeepResponse is not yet implemented")
}

func (c *Mod_md) dirMDStaplingRenewWindow(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStaplingRenewWindow directive
	return errors.New("MDStaplingRenewWindow is not yet implemented")
}

func (c *Mod_md) dirMDStoreDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStoreDir directive
	return errors.New("MDStoreDir is not yet implemented")
}

func (c *Mod_md) dirMDStoreLocks(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStoreLocks directive
	return errors.New("MDStoreLocks is not yet implemented")
}

func (c *Mod_md) dirMDWarnWindow(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDWarnWindow directive
	return errors.New("MDWarnWindow is not yet implemented")
}
