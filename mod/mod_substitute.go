// https://httpd.apache.org/docs/2.4/mod/mod_substitute.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_substitute struct{}

func (c *Mod_substitute) DirSubstitute(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Substitute directive
	return errors.New("Substitute is not yet implemented")
}

func (c *Mod_substitute) DirSubstituteInheritBefore(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SubstituteInheritBefore directive
	return errors.New("SubstituteInheritBefore is not yet implemented")
}

func (c *Mod_substitute) DirSubstituteMaxLineLength(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SubstituteMaxLineLength directive
	return errors.New("SubstituteMaxLineLength is not yet implemented")
}
