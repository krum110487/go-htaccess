// https://httpd.apache.org/docs/2.4/mod/mod_md.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_md struct{}

func (c *Mod_md) DirMDActivationDelay(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDActivationDelay directive
	return errors.New("MDActivationDelay is not yet implemented")
}

func (c *Mod_md) DirMDBaseServer(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDBaseServer directive
	return errors.New("MDBaseServer is not yet implemented")
}

func (c *Mod_md) DirMDCAChallenges(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCAChallenges directive
	return errors.New("MDCAChallenges is not yet implemented")
}

func (c *Mod_md) DirMDCertificateAgreement(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateAgreement directive
	return errors.New("MDCertificateAgreement is not yet implemented")
}

func (c *Mod_md) DirMDCertificateAuthority(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateAuthority directive
	return errors.New("MDCertificateAuthority is not yet implemented")
}

func (c *Mod_md) DirMDCertificateCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateCheck directive
	return errors.New("MDCertificateCheck is not yet implemented")
}

func (c *Mod_md) DirMDCertificateFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateFile directive
	return errors.New("MDCertificateFile is not yet implemented")
}

func (c *Mod_md) DirMDCertificateKeyFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateKeyFile directive
	return errors.New("MDCertificateKeyFile is not yet implemented")
}

func (c *Mod_md) DirMDCertificateMonitor(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateMonitor directive
	return errors.New("MDCertificateMonitor is not yet implemented")
}

func (c *Mod_md) DirMDCertificateProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateProtocol directive
	return errors.New("MDCertificateProtocol is not yet implemented")
}

func (c *Mod_md) DirMDCertificateStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDCertificateStatus directive
	return errors.New("MDCertificateStatus is not yet implemented")
}

func (c *Mod_md) DirMDChallengeDns01(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDChallengeDns01 directive
	return errors.New("MDChallengeDns01 is not yet implemented")
}

func (c *Mod_md) DirMDChallengeDns01Version(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDChallengeDns01Version directive
	return errors.New("MDChallengeDns01Version is not yet implemented")
}

func (c *Mod_md) DirMDContactEmail(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDContactEmail directive
	return errors.New("MDContactEmail is not yet implemented")
}

func (c *Mod_md) DirMDDriveMode(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDDriveMode directive
	return errors.New("MDDriveMode is not yet implemented")
}

func (c *Mod_md) DirMDExternalAccountBinding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDExternalAccountBinding directive
	return errors.New("MDExternalAccountBinding is not yet implemented")
}

func (c *Mod_md) DirMDHttpProxy(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDHttpProxy directive
	return errors.New("MDHttpProxy is not yet implemented")
}

func (c *Mod_md) DirMDMatchNames(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMatchNames directive
	return errors.New("MDMatchNames is not yet implemented")
}

func (c *Mod_md) DirMDMember(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMember directive
	return errors.New("MDMember is not yet implemented")
}

func (c *Mod_md) DirMDMembers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMembers directive
	return errors.New("MDMembers is not yet implemented")
}

func (c *Mod_md) DirMDMessageCmd(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMessageCmd directive
	return errors.New("MDMessageCmd is not yet implemented")
}

func (c *Mod_md) DirMDMustStaple(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDMustStaple directive
	return errors.New("MDMustStaple is not yet implemented")
}

func (c *Mod_md) DirMDNotifyCmd(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDNotifyCmd directive
	return errors.New("MDNotifyCmd is not yet implemented")
}

func (c *Mod_md) DirMDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDomain directive
	return errors.New("MDomain is not yet implemented")
}

func (c *Mod_md) DirMDomainSet(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <MDomainSet> directive
	return errors.New("<MDomainSet> is not yet implemented")
}

func (c *Mod_md) DirMDPortMap(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDPortMap directive
	return errors.New("MDPortMap is not yet implemented")
}

func (c *Mod_md) DirMDPrivateKeys(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDPrivateKeys directive
	return errors.New("MDPrivateKeys is not yet implemented")
}

func (c *Mod_md) DirMDRenewMode(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRenewMode directive
	return errors.New("MDRenewMode is not yet implemented")
}

func (c *Mod_md) DirMDRenewWindow(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRenewWindow directive
	return errors.New("MDRenewWindow is not yet implemented")
}

func (c *Mod_md) DirMDRequireHttps(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRequireHttps directive
	return errors.New("MDRequireHttps is not yet implemented")
}

func (c *Mod_md) DirMDRetryDelay(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRetryDelay directive
	return errors.New("MDRetryDelay is not yet implemented")
}

func (c *Mod_md) DirMDRetryFailover(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDRetryFailover directive
	return errors.New("MDRetryFailover is not yet implemented")
}

func (c *Mod_md) DirMDServerStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDServerStatus directive
	return errors.New("MDServerStatus is not yet implemented")
}

func (c *Mod_md) DirMDStapleOthers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStapleOthers directive
	return errors.New("MDStapleOthers is not yet implemented")
}

func (c *Mod_md) DirMDStapling(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStapling directive
	return errors.New("MDStapling is not yet implemented")
}

func (c *Mod_md) DirMDStaplingKeepResponse(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStaplingKeepResponse directive
	return errors.New("MDStaplingKeepResponse is not yet implemented")
}

func (c *Mod_md) DirMDStaplingRenewWindow(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStaplingRenewWindow directive
	return errors.New("MDStaplingRenewWindow is not yet implemented")
}

func (c *Mod_md) DirMDStoreDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStoreDir directive
	return errors.New("MDStoreDir is not yet implemented")
}

func (c *Mod_md) DirMDStoreLocks(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDStoreLocks directive
	return errors.New("MDStoreLocks is not yet implemented")
}

func (c *Mod_md) DirMDWarnWindow(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MDWarnWindow directive
	return errors.New("MDWarnWindow is not yet implemented")
}
