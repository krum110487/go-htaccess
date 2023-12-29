// https://httpd.apache.org/docs/2.4/mod/mod_mime.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_mime struct{}

func (Mod_mime) DirAddCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AddCharset directive
	return errors.New("AddCharset is not yet implemented")
}

func (Mod_mime) DirAddEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AddEncoding directive
	return errors.New("AddEncoding is not yet implemented")
}

func (Mod_mime) DirAddHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AddHandler directive
	return errors.New("AddHandler is not yet implemented")
}

func (Mod_mime) DirAddInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AddInputFilter directive
	return errors.New("AddInputFilter is not yet implemented")
}

func (Mod_mime) DirAddLanguage(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AddLanguage directive
	return errors.New("AddLanguage is not yet implemented")
}

func (Mod_mime) DirAddOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AddOutputFilter directive
	return errors.New("AddOutputFilter is not yet implemented")
}

func (Mod_mime) DirAddType(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: AddType directive
	return errors.New("AddType is not yet implemented")
}

func (Mod_mime) DirDefaultLanguage(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: DefaultLanguage directive
	return errors.New("DefaultLanguage is not yet implemented")
}

func (Mod_mime) DirModMimeUsePathInfo(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: ModMimeUsePathInfo directive
	return errors.New("ModMimeUsePathInfo is not yet implemented")
}

func (Mod_mime) DirMultiviewsMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: MultiviewsMatch directive
	return errors.New("MultiviewsMatch is not yet implemented")
}

func (Mod_mime) DirRemoveCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RemoveCharset directive
	return errors.New("RemoveCharset is not yet implemented")
}

func (Mod_mime) DirRemoveEncoding(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RemoveEncoding directive
	return errors.New("RemoveEncoding is not yet implemented")
}

func (Mod_mime) DirRemoveHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RemoveHandler directive
	return errors.New("RemoveHandler is not yet implemented")
}

func (Mod_mime) DirRemoveInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RemoveInputFilter directive
	return errors.New("RemoveInputFilter is not yet implemented")
}

func (Mod_mime) DirRemoveLanguage(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RemoveLanguage directive
	return errors.New("RemoveLanguage is not yet implemented")
}

func (Mod_mime) DirRemoveOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RemoveOutputFilter directive
	return errors.New("RemoveOutputFilter is not yet implemented")
}

func (Mod_mime) DirRemoveType(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: RemoveType directive
	return errors.New("RemoveType is not yet implemented")
}

func (Mod_mime) DirTypesConfig(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: TypesConfig directive
	return errors.New("TypesConfig is not yet implemented")
}
