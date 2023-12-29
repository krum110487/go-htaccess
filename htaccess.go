package htaccess

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/krum110487/go-htaccess/mod"
	"github.com/krum110487/go-htaccess/parser"
)

type Htaccess struct {
	AST            *parser.HtaccessAST
	Globals        mod.Context
	Context        mod.Context
	DirHandlers    map[string]func(string, parser.DirectiveEntry, *http.Request, *mod.Context) (bool, error)
	DirEvents      map[string]func(string, parser.DirectiveEntry, *http.Request, *mod.Context) (bool, error)
	DefaultHandler func(string, parser.DirectiveEntry, *http.Request, *mod.Context) (bool, error)
}

func New() *Htaccess {
	hta := Htaccess{}

	//Set Default Empty "Stub" handler
	hta.DefaultHandler = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *mod.Context) (bool, error) {
		return false, errors.New(fmt.Sprintf("Directive \"%s\" was called, but ignored, it has not been implemented.", s))
	}

	return &hta
}

func (hta *Htaccess) AddModules(modules ...any) {
	for _, m := range modules {
		fmt.Println(m)
	}
}

func (Htaccess) GetFunctions(interfList ...any) map[string]func(*mod.Context) error {
	flist := map[string]func(*mod.Context) error{}
	for _, x := range interfList {
		v := reflect.ValueOf(x)
		t := v.Type()
		if t.Kind() != reflect.Ptr {
			v = reflect.New(t)
			t = v.Type()
		}

		for i := 0; i < t.NumMethod(); i++ {
			name := strings.ToLower(t.Method(i).Name)
			if strings.HasPrefix(name, "dir") {
				name = name[3:]
				f := v.Method(i).Interface()
				if v, ok := f.(func(*mod.Context) error); ok {
					flist[name] = v
				}
			}
		}
	}
	return flist
}

func (hta *Htaccess) AddAllModules() {
	hta.AddModules(
		mod.Core{}, mod.Event{}, mod.Worker{}, mod.Mod_access_compat{}, mod.Mod_actions{}, mod.Mod_alias{},
		mod.Mod_allowmethods{}, mod.Mod_asis{}, mod.Mod_auth_basic{}, mod.Mod_auth_digest{}, mod.Mod_auth_form{},
		mod.Mod_authn_anon{}, mod.Mod_authn_core{}, mod.Mod_authn_dbd{}, mod.Mod_authn_dbm{}, mod.Mod_authn_file{},
		mod.Mod_authn_socache{}, mod.Mod_authnz_fcgi{}, mod.Mod_authnz_ldap{}, mod.Mod_authz_core{}, mod.Mod_authz_dbd{},
		mod.Mod_authz_host{}, mod.Mod_authz_owner{}, mod.Mod_authz_user{}, mod.Mod_autoindex{}, mod.Mod_brotli{},
		mod.Mod_buffer{}, mod.Mod_cache{}, mod.Mod_cache_disk{}, mod.Mod_cache_socache{}, mod.Mod_cern_meta{},
		mod.Mod_cgi{}, mod.Mod_cgid{}, mod.Mod_charset_lite{}, mod.Mod_data{}, mod.Mod_dav{}, mod.Mod_dav_fs{},
		mod.Mod_dav_lock{}, mod.Mod_dbd{}, mod.Mod_deflate{}, mod.Mod_dialup{}, mod.Mod_dir{}, mod.Mod_dumpio{},
		mod.Mod_echo{}, mod.Mod_env{}, mod.Mod_expires{}, mod.Mod_ext_filter{}, mod.Mod_file_cache{}, mod.Mod_filter{},
		mod.Mod_headers{}, mod.Mod_heartbeat{}, mod.Mod_heartmonitor{}, mod.Mod_http2{}, mod.Mod_ident{}, mod.Mod_imagemap{},
		mod.Mod_include{}, mod.Mod_info{}, mod.Mod_isapi{}, mod.Mod_lbmethod_bybusyness{}, mod.Mod_lbmethod_byrequests{},
		mod.Mod_lbmethod_bytraffic{}, mod.Mod_lbmethod_heartbeat{}, mod.Mod_ldap{}, mod.Mod_log_config{}, mod.Mod_log_debug{},
		mod.Mod_log_forensic{}, mod.Mod_logio{}, mod.Mod_lua{}, mod.Mod_macro{}, mod.Mod_md{}, mod.Mod_mime{},
		mod.Mod_mime_magic{}, mod.Mod_negotiation{}, mod.Mod_nw_ssl{}, mod.Mod_privileges{}, mod.Mod_privileges{},
		mod.Mod_proxy{}, mod.Mod_proxy_ajp{}, mod.Mod_proxy_balancer{}, mod.Mod_proxy_connect{}, mod.Mod_proxy_express{},
		mod.Mod_proxy_fcgi{}, mod.Mod_proxy_fdpass{}, mod.Mod_proxy_ftp{}, mod.Mod_proxy_hcheck{}, mod.Mod_proxy_html{},
		mod.Mod_proxy_http{}, mod.Mod_proxy_http2{}, mod.Mod_proxy_scgi{}, mod.Mod_proxy_uwsgi{}, mod.Mod_proxy_wstunnel{},
		mod.Mod_ratelimit{}, mod.Mod_reflector{}, mod.Mod_remoteip{}, mod.Mod_reqtimeout{}, mod.Mod_request{},
		mod.Mod_rewrite{}, mod.Mod_sed{}, mod.Mod_session{}, mod.Mod_session_cookie{}, mod.Mod_session_crypto{},
		mod.Mod_session_dbd{}, mod.Mod_setenvif{}, mod.Mod_slotmem_plain{}, mod.Mod_slotmem_shm{}, mod.Mod_so{},
		mod.Mod_socache_dbm{}, mod.Mod_socache_dc{}, mod.Mod_socache_memcache{}, mod.Mod_socache_redis{},
		mod.Mod_socache_shmcb{}, mod.Mod_speling{}, mod.Mod_ssl{}, mod.Mod_status{}, mod.Mod_substitute{},
		mod.Mod_suexec{}, mod.Mod_systemd{}, mod.Mod_tls{}, mod.Mod_unique_id{}, mod.Mod_unixd{}, mod.Mod_userdir{},
		mod.Mod_usertrack{}, mod.Mod_version{}, mod.Mod_vhost_alias{}, mod.Mod_watchdog{}, mod.Mod_xml2enc{},
		mod.Mpmt_os2{}, mod.Mpm_common{}, mod.Mpm_netware{}, mod.Mpm_winnt{}, mod.Prefork{}, mod.Worker{},
	)
}

func (hta *Htaccess) AddHandler(directives []string, handler func(string, parser.DirectiveEntry, *http.Request, *mod.Context) (bool, error)) {
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
	hta.Context = mod.Context{}

	hta.DefaultHandler = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *mod.Context) (bool, error) {
		return false, errors.New(fmt.Sprintf("Directive \"%s\" was called, but ignored, it has not been implemented.", s))
	}

	loadDefaultHandlers(hta.DirHandlers)

	return hta, nil
}

func loadDefaultHandlers(hdlrMap map[string]func(string, parser.DirectiveEntry, *http.Request, *mod.Context) (bool, error)) {
	hdlrMap["if"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *mod.Context) (bool, error) {
		return de.IfGroup.Expression.Resolve()
	}
	hdlrMap["elseif"] = hdlrMap["if"]

	hdlrMap["RewriteEngine"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *mod.Context) (bool, error) {
		if err := paramCheckPass(de, 1, 1); err != nil {
			return false, err
		}

		//ctx.requestConfig.ModRewrite.EngineOn = false
		//if strings.ToLower(de.Params[0]) == "on" {
		//	ctx.requestConfig.ModRewrite.EngineOn = true
		//}

		return true, nil
	}

	hdlrMap["RewriteCond"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *mod.Context) (bool, error) {
		//TODO: Should this error?
		//if !ctx.requestConfig.ModRewrite.EngineOn {
		//	return true, nil
		//}

		if err := paramCheckPass(de, 2, 2); err != nil {
			return false, err
		}

		//TODO: Do the RewriteCond logic check and set the results before putting it on the stack group.
		de.Results = true

		//ctx.stackGroup["RewriteCond"] = append(ctx.stackGroup["RewriteCond"], de)
		return true, nil
	}

	hdlrMap["RewriteRule"] = func(s string, de parser.DirectiveEntry, req *http.Request, ctx *mod.Context) (bool, error) {
		//Grab the conditions, then empty the stackGroup
		//conditions := ctx.stackGroup["RewriteCond"]
		//ctx.stackGroup["RewriteCond"] = []parser.DirectiveEntry{}

		//if !ctx.requestConfig.ModRewrite.EngineOn {
		//	return true, nil
		//}

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

/*
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
*/

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
