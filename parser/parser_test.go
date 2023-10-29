package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"testing"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func testMatch(t *testing.T, tests [][]string) {
	failed := false
	for i, test := range tests {
		name := fmt.Sprintf("Test %d", i)
		src := test[0]
		res := fmt.Sprintf(`%s`, test[1])

		hta, err := htaccessParser.ParseString(name, src)
		if err != nil {
			panic(err)
		}

		data, _ := json.Marshal(hta)
		jsonStr := string(data)
		if jsonStr != res {
			failed = true
			t.Logf("Test %d failed to match:\n\tSource:    %s\n\tParsedStr: %s\n\tExpected:  %s",
				i+1, src, jsonStr, res)
		}
	}

	if failed {
		t.FailNow()
	}
}

func TestParserRewriteRule(t *testing.T) {
	tests := [][]string{
		{`RewriteRule "^search/(.*)$" "/search.php?term=$1" "[B= ?]"`,
			`{"directives":[{"IfGroup":null,"Tag":null,"Name":"RewriteRule","Params":[{"rawString":"^search/(.*)$"},{"rawString":"/search.php?term=$1","variableName":"/search.php?term","assignValue":"$1"},{"rawString":"[B= ?]","flags":{"flagMap":{"B":{"valueStr":" ?"}}}}],"Results":false}]}`},
		{`RewriteRule "^search/(.*)$" "/search.php?term=$1" [B=?]`,
			`{"directives":[{"IfGroup":null,"Tag":null,"Name":"RewriteRule","Params":[{"rawString":"^search/(.*)$"},{"rawString":"/search.php?term=$1","variableName":"/search.php?term","assignValue":"$1"},{"rawString":"[B=?]","flags":{"flagMap":{"B":{"valueStr":"?"}}}}],"Results":false}]}`},
	}

	testMatch(t, tests)
}

func TestParserRewriteCond(t *testing.T) {
	tests := [][]string{
		{`RewriteCond "%{HTTP_HOST}" "(.*)"
		  RewriteRule "^/(.*)"       "/sites/%1/$1"`,
			`{"directives":[{"IfGroup":null,"Tag":null,"Name":"RewriteCond","Params":[{"rawString":"%{HTTP_HOST}","variableName":"HTTP_HOST"},{"rawString":"(.*)"}],"Results":false},{"IfGroup":null,"Tag":null,"Name":"RewriteRule","Params":[{"rawString":"^/(.*)"},{"rawString":"/sites/%1/$1"}],"Results":false}]}`},
		{`RewriteCond "%{QUERY_STRING}" "hack"
		  RewriteCond "%{HTTP_COOKIE}"  !go
		  RewriteRule "."               "-"   [F]`,
			`{"directives":[{"IfGroup":null,"Tag":null,"Name":"RewriteCond","Params":[{"rawString":"%{QUERY_STRING}","variableName":"QUERY_STRING"},{"rawString":"hack"}],"Results":false},{"IfGroup":null,"Tag":null,"Name":"RewriteCond","Params":[{"rawString":"%{HTTP_COOKIE}","variableName":"HTTP_COOKIE"},{"rawString":"!go","notOp":true,"notOpString":"go"}],"Results":false},{"IfGroup":null,"Tag":null,"Name":"RewriteRule","Params":[{"rawString":"."},{"rawString":"-"},{"rawString":"[F]","flags":{"flagMap":{"F":{"valueBool":true}}}}],"Results":false}]}`},
	}

	testMatch(t, tests)
}

func TestParserSetEnvIf(t *testing.T) {
	tests := [][]string{
		{`SetEnvIf user-agent "WebBandit" stayout=1`,
			`{"directives":[{"IfGroup":null,"Tag":null,"Name":"SetEnvIf","Params":[{"rawString":"user-agent"},{"rawString":"WebBandit"},{"rawString":"stayout=1","variableName":"stayout","assignValue":"1"}],"Results":false}]}`},
	}

	testMatch(t, tests)
}

func TestParserDirectoryIndex(t *testing.T) {
	tests := [][]string{
		{`RewriteEngine On
		  DirectoryIndex index.php index.html`,
			`{"directives":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"On"}],"Results":false},{"IfGroup":null,"Tag":null,"Name":"DirectoryIndex","Params":[{"rawString":"index.php"},{"rawString":"index.html"}],"Results":false}]}`},
		{`Require all granted
		  Require not env stayout`,
			`{"directives":[{"IfGroup":null,"Tag":null,"Name":"Require","Params":[{"rawString":"all"},{"rawString":"granted"}],"Results":false},{"IfGroup":null,"Tag":null,"Name":"Require","Params":[{"rawString":"not"},{"rawString":"env"},{"rawString":"stayout"}],"Results":false}]}`},
	}

	testMatch(t, tests)
}

func TestParserIfThenElse(t *testing.T) {
	tests := [][]string{
		{`<If "-R '10.1.0.0/16'">
			RewriteEngine On
		  </If>
		  <ElseIf "-R '10.2.0.0/8'">
		    # Do Nothing
		  </ElseIf>
		  <ElseIf "-R '10.3.0.0/8'">
		    RewriteEngine Off
		  </ElseIf>
		  <Else>
		    RewriteEngine On
		  </Else>`,
			`{"directives":[{"IfGroup":{"Name":"If","Expression":{"operator":"-R","rightArgument":"'10.1.0.0/16'"},"IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"On"}],"Results":false}],"Tags":null,"TagClose":"If","ElseIf":[{"Name":"ElseIf","Expression":{"operator":"-R","rightArgument":"'10.2.0.0/8'"},"IfGroup":null,"Dirs":null,"Tags":null,"TagClose":"ElseIf"},{"Name":"ElseIf","Expression":{"operator":"-R","rightArgument":"'10.3.0.0/8'"},"IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"Off"}],"Results":false}],"Tags":null,"TagClose":"ElseIf"}],"Else":{"Name":"Else","IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"On"}],"Results":false}],"Tags":null,"TagClose":"Else"}},"Tag":null,"Name":"","Params":null,"Results":false}]}`},
	}

	testMatch(t, tests)
}

func TestParserNestedIfThenElse(t *testing.T) {
	tests := [][]string{
		{`<If "-R '10.1.0.0/16'">
		    <If "-R '10.1.0.0/16'">
		      RewriteEngine On
	        </If>
			<Else>
			  RewriteEngine Off
			</Else>
		  </If>
		  <ElseIf "-R '10.2.0.0/8'">
		    <If "-R '10.1.0.0/16'">
		      RewriteEngine On
		    </If>
			<Else>
			  RewriteEngine Off
			</Else>
		  </ElseIf>
		  <Else>
		    <If "-R '10.1.0.0/16'">
		      RewriteEngine On
		    </If>
		    <Else>
		      RewriteEngine Off
		    </Else>
		  </Else>`,
			`{"directives":[{"IfGroup":{"Name":"If","Expression":{"operator":"-R","rightArgument":"'10.1.0.0/16'"},"IfGroup":{"Name":"If","Expression":{"operator":"-R","rightArgument":"'10.1.0.0/16'"},"IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"On"}],"Results":false}],"Tags":null,"TagClose":"If","ElseIf":null,"Else":{"Name":"Else","IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"Off"}],"Results":false}],"Tags":null,"TagClose":"Else"}},"Dirs":null,"Tags":null,"TagClose":"If","ElseIf":[{"Name":"ElseIf","Expression":{"operator":"-R","rightArgument":"'10.2.0.0/8'"},"IfGroup":{"Name":"If","Expression":{"operator":"-R","rightArgument":"'10.1.0.0/16'"},"IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"On"}],"Results":false}],"Tags":null,"TagClose":"If","ElseIf":null,"Else":{"Name":"Else","IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"Off"}],"Results":false}],"Tags":null,"TagClose":"Else"}},"Dirs":null,"Tags":null,"TagClose":"ElseIf"}],"Else":{"Name":"Else","IfGroup":{"Name":"If","Expression":{"operator":"-R","rightArgument":"'10.1.0.0/16'"},"IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"On"}],"Results":false}],"Tags":null,"TagClose":"If","ElseIf":null,"Else":{"Name":"Else","IfGroup":null,"Dirs":[{"IfGroup":null,"Tag":null,"Name":"RewriteEngine","Params":[{"rawString":"Off"}],"Results":false}],"Tags":null,"TagClose":"Else"}},"Dirs":null,"Tags":null,"TagClose":"Else"}},"Tag":null,"Name":"","Params":null,"Results":false}]}`},
	}

	testMatch(t, tests)
}

func TestParserInvalidDirective(t *testing.T) {

}

func TestParserFiles(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../testFiles")

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		path := dir + "/" + f.Name()
		_, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := os.ReadFile(path)
		_, err = htaccessParser.ParseBytes(f.Name(), b)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
	}
}
