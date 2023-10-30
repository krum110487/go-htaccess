//nolint:gosec
package main

import (
	"encoding/json"
	"fmt"

	"github.com/krum110487/go-htaccess/parser"
)

func main() {
	//Parse to AST
	//Make a second p
	v, err := parser.Parse(
		`
RewriteCond %{HTTP_USER_AGENT} "=This Robot/1.0"
RewriteRule test2

RewriteCond /var/www/%{REQUEST_URI} !-f
RewriteRule ^(.+) /other/archive/$1 [R]

RewriteCond expr "! %{HTTP_REFERER} -strmatch '*://%{HTTP_HOST}/*'"
RewriteRule "^/images" "-" [F]

RewriteCond "%{REMOTE_HOST}"  "^host1"  [OR]
RewriteCond "%{REMOTE_HOST}"  "^host2"  [OR,E=VAR:VALUE]
RewriteCond "%{REMOTE_HOST}"  "^host3"
RewriteRule test
		`)

	//RewriteCond pushes to rewritecond set
	//expr pushes to expression set.

	if err != nil {
		panic(err)
	}

	//repr.Println(v, repr.Indent("  "), repr.OmitEmpty(true))

	d, err := json.MarshalIndent(v, "  ", "    ")
	fmt.Println(string(d))
	//fmt.Printf("%d", v.Directives[len(v.Directives)-1].Entry.Paramaters[0].Number)
}
