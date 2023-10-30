// https://httpd.apache.org/docs/2.4/mod/mod_cern_meta.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_cern_meta struct{}

func (c *Mod_cern_meta) dirMetaDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MetaDir directive
	return errors.New("MetaDir is not yet implemented")
}

func (c *Mod_cern_meta) dirMetaFiles(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MetaFiles directive
	return errors.New("MetaFiles is not yet implemented")
}

func (c *Mod_cern_meta) dirMetaSuffix(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MetaSuffix directive
	return errors.New("MetaSuffix is not yet implemented")
}
