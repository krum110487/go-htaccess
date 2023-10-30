// https://httpd.apache.org/docs/2.4/mod/mod_auth_digest.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_auth_digest struct{}

func (c *Mod_auth_digest) DirAuthDigestAlgorithm(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDigestAlgorithm directive
	return errors.New("AuthDigestAlgorithm is not yet implemented")
}

func (c *Mod_auth_digest) DirAuthDigestDomain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDigestDomain directive
	return errors.New("AuthDigestDomain is not yet implemented")
}

func (c *Mod_auth_digest) DirAuthDigestNonceLifetime(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDigestNonceLifetime directive
	return errors.New("AuthDigestNonceLifetime is not yet implemented")
}

func (c *Mod_auth_digest) DirAuthDigestProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDigestProvider directive
	return errors.New("AuthDigestProvider is not yet implemented")
}

func (c *Mod_auth_digest) DirAuthDigestQop(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDigestQop directive
	return errors.New("AuthDigestQop is not yet implemented")
}

func (c *Mod_auth_digest) DirAuthDigestShmemSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AuthDigestShmemSize directive
	return errors.New("AuthDigestShmemSize is not yet implemented")
}
