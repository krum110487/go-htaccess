// https://httpd.apache.org/docs/2.4/mod/mod_autoindex.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_autoindex struct{}

func (c *Mod_autoindex) DirAddAlt(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddAlt directive
	return errors.New("AddAlt is not yet implemented")
}

func (c *Mod_autoindex) DirAddAltByEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddAltByEncoding directive
	return errors.New("AddAltByEncoding is not yet implemented")
}

func (c *Mod_autoindex) DirAddAltByType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddAltByType directive
	return errors.New("AddAltByType is not yet implemented")
}

func (c *Mod_autoindex) DirAddDescription(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddDescription directive
	return errors.New("AddDescription is not yet implemented")
}

func (c *Mod_autoindex) DirAddIcon(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddIcon directive
	return errors.New("AddIcon is not yet implemented")
}

func (c *Mod_autoindex) DirAddIconByEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddIconByEncoding directive
	return errors.New("AddIconByEncoding is not yet implemented")
}

func (c *Mod_autoindex) DirAddIconByType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddIconByType directive
	return errors.New("AddIconByType is not yet implemented")
}

func (c *Mod_autoindex) DirDefaultIcon(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DefaultIcon directive
	return errors.New("DefaultIcon is not yet implemented")
}

func (c *Mod_autoindex) DirHeaderName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HeaderName directive
	return errors.New("HeaderName is not yet implemented")
}

func (c *Mod_autoindex) DirIndexHeadInsert(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexHeadInsert directive
	return errors.New("IndexHeadInsert is not yet implemented")
}

func (c *Mod_autoindex) DirIndexIgnore(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexIgnore directive
	return errors.New("IndexIgnore is not yet implemented")
}

func (c *Mod_autoindex) DirIndexIgnoreReset(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexIgnoreReset directive
	return errors.New("IndexIgnoreReset is not yet implemented")
}

func (c *Mod_autoindex) DirIndexOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexOptions directive
	return errors.New("IndexOptions is not yet implemented")
}

func (c *Mod_autoindex) DirIndexOrderDefault(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexOrderDefault directive
	return errors.New("IndexOrderDefault is not yet implemented")
}

func (c *Mod_autoindex) DirIndexStyleSheet(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IndexStyleSheet directive
	return errors.New("IndexStyleSheet is not yet implemented")
}

func (c *Mod_autoindex) DirReadmeName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ReadmeName directive
	return errors.New("ReadmeName is not yet implemented")
}
