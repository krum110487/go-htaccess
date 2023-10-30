// https://httpd.apache.org/docs/2.4/mod/mod_autoindex.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_autoindex struct{}

func (c *Mod_autoindex) dirAddAlt(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddAlt directive
	return errors.New("AddAlt is not yet implemented")
}

func (c *Mod_autoindex) dirAddAltByEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddAltByEncoding directive
	return errors.New("AddAltByEncoding is not yet implemented")
}

func (c *Mod_autoindex) dirAddAltByType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddAltByType directive
	return errors.New("AddAltByType is not yet implemented")
}

func (c *Mod_autoindex) dirAddDescription(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddDescription directive
	return errors.New("AddDescription is not yet implemented")
}

func (c *Mod_autoindex) dirAddIcon(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddIcon directive
	return errors.New("AddIcon is not yet implemented")
}

func (c *Mod_autoindex) dirAddIconByEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddIconByEncoding directive
	return errors.New("AddIconByEncoding is not yet implemented")
}

func (c *Mod_autoindex) dirAddIconByType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddIconByType directive
	return errors.New("AddIconByType is not yet implemented")
}

func (c *Mod_autoindex) dirDefaultIcon(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DefaultIcon directive
	return errors.New("DefaultIcon is not yet implemented")
}

func (c *Mod_autoindex) dirHeaderName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HeaderName directive
	return errors.New("HeaderName is not yet implemented")
}

func (c *Mod_autoindex) dirIndexHeadInsert(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexHeadInsert directive
	return errors.New("IndexHeadInsert is not yet implemented")
}

func (c *Mod_autoindex) dirIndexIgnore(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexIgnore directive
	return errors.New("IndexIgnore is not yet implemented")
}

func (c *Mod_autoindex) dirIndexIgnoreReset(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexIgnoreReset directive
	return errors.New("IndexIgnoreReset is not yet implemented")
}

func (c *Mod_autoindex) dirIndexOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexOptions directive
	return errors.New("IndexOptions is not yet implemented")
}

func (c *Mod_autoindex) dirIndexOrderDefault(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexOrderDefault directive
	return errors.New("IndexOrderDefault is not yet implemented")
}

func (c *Mod_autoindex) dirIndexStyleSheet(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexStyleSheet directive
	return errors.New("IndexStyleSheet is not yet implemented")
}

func (c *Mod_autoindex) dirReadmeName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ReadmeName directive
	return errors.New("ReadmeName is not yet implemented")
}
