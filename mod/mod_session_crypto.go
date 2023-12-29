// https://httpd.apache.org/docs/2.4/mod/mod_session_crypto.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_session_crypto struct{}

func (Mod_session_crypto) DirSessionCryptoCipher(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionCryptoCipher directive
	return errors.New("SessionCryptoCipher is not yet implemented")
}

func (Mod_session_crypto) DirSessionCryptoDriver(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionCryptoDriver directive
	return errors.New("SessionCryptoDriver is not yet implemented")
}

func (Mod_session_crypto) DirSessionCryptoPassphrase(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionCryptoPassphrase directive
	return errors.New("SessionCryptoPassphrase is not yet implemented")
}

func (Mod_session_crypto) DirSessionCryptoPassphraseFile(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: SessionCryptoPassphraseFile directive
	return errors.New("SessionCryptoPassphraseFile is not yet implemented")
}
