// https://httpd.apache.org/docs/2.4/mod/mod_xml2enc.html
package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess/parser"
)

type Mod_xml2enc struct{}

func (Mod_xml2enc) Dirxml2EncAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: xml2EncAlias directive
	return errors.New("xml2EncAlias is not yet implemented")
}

func (Mod_xml2enc) Dirxml2EncDefault(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: xml2EncDefault directive
	return errors.New("xml2EncDefault is not yet implemented")
}

func (Mod_xml2enc) Dirxml2StartParse(dir parser.DirectiveEntry, req *http.Request, runCtx *Context) error {
	//TODO: xml2StartParse directive
	return errors.New("xml2StartParse is not yet implemented")
}
