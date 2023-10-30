package htaccess

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/krum110487/go-htaccess/parser"
)

type Htaccess struct {
	AST            *parser.HtaccessAST
	Globals        Context
	Context        Context
	DirHandlers    map[string]func(string, parser.DirectiveEntry, *http.Request, *Context) (bool, error)
	DirEvents      map[string]func(string, parser.DirectiveEntry, *http.Request, *Context) (bool, error)
	DefaultHandler func(string, parser.DirectiveEntry, *http.Request, *Context) (bool, error)
}

func GetParserWithDefaultHandlers() Htaccess {
	hta := Htaccess{}

	//Set Default Empty "Stub" handler
	hta.DefaultHandler = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *Context) (bool, error) {
		return false, errors.New(fmt.Sprintf("Directive \"%s\" was called, but ignored, it has not been implemented.", s))
	}

	return hta
}

func (hta *Htaccess) AddHandler(directives []string, handler func(string, parser.DirectiveEntry, *http.Request, *Context) (bool, error)) {
	for _, dir := range directives {
		hta.DirHandlers[dir] = handler
	}
}

func (hta *Htaccess) DelHandler(directives []string) {
	for _, dir := range directives {
		if _, ok := hta.DirHandlers[dir]; ok {
			delete(hta.DirHandlers, dir)
		}
	}
}

func ParseString(fileContents string) (Htaccess, error) {
	var err error = nil

	hta := Htaccess{}
	hta.AST, err = parser.Parse(fileContents)
	if err != nil {
		return Htaccess{}, err
	}
	hta.Context = Context{}

	hta.DefaultHandler = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *Context) (bool, error) {
		return false, errors.New(fmt.Sprintf("Directive \"%s\" was called, but ignored, it has not been implemented.", s))
	}

	loadDefaultHandlers(hta.DirHandlers)

	return hta, nil
}

func loadDefaultHandlers(hdlrMap map[string]func(string, parser.DirectiveEntry, *http.Request, *Context) (bool, error)) {
	hdlrMap["if"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *Context) (bool, error) {
		return de.IfGroup.Expression.Resolve()
	}
	hdlrMap["elseif"] = hdlrMap["if"]

	hdlrMap["RewriteEngine"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *Context) (bool, error) {
		if err := paramCheckPass(de, 1, 1); err != nil {
			return false, err
		}

		ctx.requestConfig.ModRewrite.EngineOn = false
		//if strings.ToLower(de.Params[0]) == "on" {
		//	ctx.requestConfig.ModRewrite.EngineOn = true
		//}

		return true, nil
	}

	hdlrMap["RewriteCond"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *Context) (bool, error) {
		//TODO: Should this error?
		if !ctx.requestConfig.ModRewrite.EngineOn {
			return true, nil
		}

		if err := paramCheckPass(de, 2, 2); err != nil {
			return false, err
		}

		//TODO: Do the RewriteCond logic check and set the results before putting it on the stack group.
		de.Results = true

		ctx.stackGroup["RewriteCond"] = append(ctx.stackGroup["RewriteCond"], de)
		return true, nil
	}

	hdlrMap["RewriteRule"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *Context) (bool, error) {
		//Grab the conditions, then empty the stackGroup
		//conditions := ctx.stackGroup["RewriteCond"]
		ctx.stackGroup["RewriteCond"] = []parser.DirectiveEntry{}

		if !ctx.requestConfig.ModRewrite.EngineOn {
			return true, nil
		}

		if err := paramCheckPass(de, 2, 2); err != nil {
			return false, err
		}

		/*
			res, err := testRewriteConds(conditions, req)
			if err != nil {
				return false, err
			}
		*/

		//de.Results = res
		return true, nil
	}
}

/*
func testRewriteConds(conds []parser.DirectiveEntry, req *http.Request) (bool, error) {
	result := false
	for i := len(conds) - 1; i >= 0; i-- {
		cond := conds[i].RewriteCond
		match, err := testPattern(cond, req)
		if err != nil {
			return false, err
		}

		if i == len(conds)-1 {
			result = match
		} else {
			if cond.Flags.Exists("OR") {
				result = result || match
			} else {
				result = result && match
			}
		}
	}
	return result, nil
}
*/

func testPattern(cond *parser.RWCondEntry, req *http.Request) (bool, error) {
	testStr := replaceVariables(req, cond.TestString)
	ptrnRightArg := cond.Pattern.RightArgument

	switch cond.Pattern.Operator {
	case "==":
		return testStr == ptrnRightArg, nil
	case "!=":
		return testStr != ptrnRightArg, nil
	case "<":
		return testStr < ptrnRightArg, nil
	case "<=":
		return testStr <= ptrnRightArg, nil
	case ">":
		return testStr > ptrnRightArg, nil
	case ">=":
		return testStr >= ptrnRightArg, nil
	case "=~":
		regex := regexp.MustCompile(ptrnRightArg)
		return regex.MatchString(testStr), nil
	default:
		return false, errors.New(fmt.Sprintf("Unhandeled operator %s", cond.Pattern.Operator))
	}
}

func replaceVariables(request *http.Request, str string) string {
	// Create a map of the available variables and their corresponding values
	variables := map[string]string{
		"%{API_VERSION}":           request.Proto,
		"%{AUTH_TYPE}":             request.Header.Get("Authorization"),
		"%{CONTENT_LENGTH}":        strconv.FormatInt(request.ContentLength, 10),
		"%{CONTENT_TYPE}":          request.Header.Get("Content-Type"),
		"%{DOCUMENT_ROOT}":         "/var/www/html", // example value, change to your actual document root
		"%{GATEWAY_INTERFACE}":     "CGI/1.1",       // example value, change to your actual gateway interface
		"%{HTTP_ACCEPT}":           request.Header.Get("Accept"),
		"%{HTTP_ACCEPT_CHARSET}":   request.Header.Get("Accept-Charset"),
		"%{HTTP_ACCEPT_ENCODING}":  request.Header.Get("Accept-Encoding"),
		"%{HTTP_ACCEPT_LANGUAGE}":  request.Header.Get("Accept-Language"),
		"%{HTTP_CACHE_CONTROL}":    request.Header.Get("Cache-Control"),
		"%{HTTP_CONNECTION}":       request.Header.Get("Connection"),
		"%{HTTP_COOKIE}":           request.Header.Get("Cookie"),
		"%{HTTP_FORWARDED}":        request.Header.Get("Forwarded"),
		"%{HTTP_HOST}":             request.Host,
		"%{HTTP_KEEP_ALIVE}":       request.Header.Get("Keep-Alive"),
		"%{HTTP_PROXY_CONNECTION}": request.Header.Get("Proxy-Connection"),
		"%{HTTP_REFERER}":          request.Header.Get("Referer"),
		"%{HTTP_USER_AGENT}":       request.Header.Get("User-Agent"),
		"%{HTTPS}":                 strconv.FormatBool(request.TLS != nil),
		"%{SSL_CIPHER}":            "", // example value, change to your actual SSL cipher
		"%{SSL_CIPHER_ALGKEYSIZE}": "", // example value, change to your actual SSL cipher algkeysize
		"%{SSL_CIPHER_EXPORT}":     "", // example value, change to your actual SSL cipher export
		"%{SSL_CIPHER_USEKEYSIZE}": "", // example value, change to your actual SSL cipher usekeysize
		"%{SSL_CLIENT_VERIFY}":     "", // example value, change to your actual SSL client verify
		"%{SSL_PROTOCOL}":          "", // example value, change to your actual SSL protocol
		"%{SSL_SERVER_A_KEY}":      "", // example value, change to your actual SSL server A key
		"%{SSL_SERVER_A_SIG}":      "", // example value, change to your actual SSL server A signature
		"%{SSL_SERVER_CERT}":       "", // example value, change to your actual SSL server certificate
		"%{SSL_SERVER_I_DN}":       "", // example value, change to your actual SSL server issuer DN
		"%{SSL_SERVER_I_DN_C}":     "", // example value, change to your actual SSL server issuer DN country
		"%{SSL_SERVER_I_DN_CN}":    "", // example value, change to your actual SSL server issuer DN common name
		"%{SSL_SERVER_I_DN_L}":     "", // example value, change to your actual SSL server issuer DN locality
		"%{SSL_SERVER_I_DN_O}":     "", // example value, change to your actual SSL server issuer DN organization
		"%{SSL_SERVER_I_DN_OU}":    "", // example value, change to your actual SSL server issuer DN organizational unit
		"%{SSL_SERVER_I_DN_ST}":    "", // example value, change to your actual SSL server issuer DN state
		"%{SSL_SERVER_M_SERIAL":    "",
		"%{SSL_SERVER_M_VERSION":   "",
		"%{SSL_SERVER_S_DN":        "",
		"%{SSL_SERVER_S_DN_CN":     "",
		"%{SSL_SERVER_S_DN_O":      "",
		"%{SSL_SERVER_S_DN_OU":     "",
		"%{SSL_SERVER_V_END":       "",
		"%{SSL_SERVER_V_START":     "",
		"%{SSL_SESSION_ID":         "",
		"%{SSL_VERSION_INTERFACE":  "",
		"%{SSL_VERSION_LIBRARY":    "",
	}

	for key, val := range variables {
		str = strings.ReplaceAll(str, key, val)
	}

	return str
}

func paramCheckPass(de parser.DirectiveEntry, min int, max int) error {
	if de.Params == nil {
		return errors.New(fmt.Sprintf("There are no Parameters the directive \"%s\" which requires at least %d parameters.", de.Name, min))
	}

	if len(de.Params) < min {
		return errors.New(fmt.Sprintf("There are not enough Parameters for the directive \"%s\" which requires at least %d parameters.", de.Name, min))
	}

	if len(de.Params) > max {
		return errors.New(fmt.Sprintf("There too many Parameters for the directive \"%s\" which requires at most %d parameters.", de.Name, max))
	}

	return nil
}

/*
reqCtx = Context.NewContext{}
Split the Request into A Full path to each folder including root "/"

Loop through each Path on the Rquest
  htaAST = hta[urlPath]
  if htaAST == nil then
    break

  runCtx = Context.NewContext{}
  htaAST.Process(reqCtx, runCtx)



Process
  Loop through all directives and resolve them
  If a Tag is found, check the expression, if it resolves, then run those directives too
  Any directive can store to the Contexts as needed.

*/
