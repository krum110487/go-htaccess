// https://httpd.apache.org/docs/2.4/mod/mod_session_crypto.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_session_crypto struct{}

func (c *Mod_session_crypto) dirSessionCryptoCipher(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionCryptoCipher directive
	return errors.New("SessionCryptoCipher is not yet implemented")
}

func (c *Mod_session_crypto) dirSessionCryptoDriver(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionCryptoDriver directive
	return errors.New("SessionCryptoDriver is not yet implemented")
}

func (c *Mod_session_crypto) dirSessionCryptoPassphrase(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionCryptoPassphrase directive
	return errors.New("SessionCryptoPassphrase is not yet implemented")
}

func (c *Mod_session_crypto) dirSessionCryptoPassphraseFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SessionCryptoPassphraseFile directive
	return errors.New("SessionCryptoPassphraseFile is not yet implemented")
}
