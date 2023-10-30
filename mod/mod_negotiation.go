// https://httpd.apache.org/docs/2.4/mod/mod_negotiation.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_negotiation struct{}

func (c *Mod_negotiation) DirCacheNegotiatedDocs(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CacheNegotiatedDocs directive
	return errors.New("CacheNegotiatedDocs is not yet implemented")
}

func (c *Mod_negotiation) DirForceLanguagePriority(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ForceLanguagePriority directive
	return errors.New("ForceLanguagePriority is not yet implemented")
}

func (c *Mod_negotiation) DirLanguagePriority(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LanguagePriority directive
	return errors.New("LanguagePriority is not yet implemented")
}
