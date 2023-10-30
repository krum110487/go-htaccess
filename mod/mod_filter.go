// https://httpd.apache.org/docs/2.4/mod/mod_filter.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_filter struct{}

func (c *Mod_filter) DirAddOutputFilterByType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddOutputFilterByType directive
	return errors.New("AddOutputFilterByType is not yet implemented")
}

func (c *Mod_filter) DirFilterChain(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FilterChain directive
	return errors.New("FilterChain is not yet implemented")
}

func (c *Mod_filter) DirFilterDeclare(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FilterDeclare directive
	return errors.New("FilterDeclare is not yet implemented")
}

func (c *Mod_filter) DirFilterProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FilterProtocol directive
	return errors.New("FilterProtocol is not yet implemented")
}

func (c *Mod_filter) DirFilterProvider(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FilterProvider directive
	return errors.New("FilterProvider is not yet implemented")
}

func (c *Mod_filter) DirFilterTrace(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FilterTrace directive
	return errors.New("FilterTrace is not yet implemented")
}
