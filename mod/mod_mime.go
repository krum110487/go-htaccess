// https://httpd.apache.org/docs/2.4/mod/mod_mime.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Mod_mime struct{}

func (c *Mod_mime) dirAddCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddCharset directive
	return errors.New("AddCharset is not yet implemented")
}

func (c *Mod_mime) dirAddEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddEncoding directive
	return errors.New("AddEncoding is not yet implemented")
}

func (c *Mod_mime) dirAddHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddHandler directive
	return errors.New("AddHandler is not yet implemented")
}

func (c *Mod_mime) dirAddInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddInputFilter directive
	return errors.New("AddInputFilter is not yet implemented")
}

func (c *Mod_mime) dirAddLanguage(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddLanguage directive
	return errors.New("AddLanguage is not yet implemented")
}

func (c *Mod_mime) dirAddOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddOutputFilter directive
	return errors.New("AddOutputFilter is not yet implemented")
}

func (c *Mod_mime) dirAddType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddType directive
	return errors.New("AddType is not yet implemented")
}

func (c *Mod_mime) dirDefaultLanguage(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DefaultLanguage directive
	return errors.New("DefaultLanguage is not yet implemented")
}

func (c *Mod_mime) dirModMimeUsePathInfo(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ModMimeUsePathInfo directive
	return errors.New("ModMimeUsePathInfo is not yet implemented")
}

func (c *Mod_mime) dirMultiviewsMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MultiviewsMatch directive
	return errors.New("MultiviewsMatch is not yet implemented")
}

func (c *Mod_mime) dirRemoveCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoveCharset directive
	return errors.New("RemoveCharset is not yet implemented")
}

func (c *Mod_mime) dirRemoveEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoveEncoding directive
	return errors.New("RemoveEncoding is not yet implemented")
}

func (c *Mod_mime) dirRemoveHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoveHandler directive
	return errors.New("RemoveHandler is not yet implemented")
}

func (c *Mod_mime) dirRemoveInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoveInputFilter directive
	return errors.New("RemoveInputFilter is not yet implemented")
}

func (c *Mod_mime) dirRemoveLanguage(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoveLanguage directive
	return errors.New("RemoveLanguage is not yet implemented")
}

func (c *Mod_mime) dirRemoveOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoveOutputFilter directive
	return errors.New("RemoveOutputFilter is not yet implemented")
}

func (c *Mod_mime) dirRemoveType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RemoveType directive
	return errors.New("RemoveType is not yet implemented")
}

func (c *Mod_mime) dirTypesConfig(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TypesConfig directive
	return errors.New("TypesConfig is not yet implemented")
}
