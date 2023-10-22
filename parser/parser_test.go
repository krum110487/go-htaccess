package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
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
		res := fmt.Sprintf(`{"directives":[%s]}`, test[1])

		hta, err := htaccessParser.ParseString(name, src)
		if err != nil {
			panic(err)
		}

		for _, dir := range hta.Dirs {
			dir.Pos = lexer.Position{}
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

func generateDirectiveString(name string, flag string, params ...string) string {
	rwRule := DirectiveEntry{}

	rwRule.Name = name
	if flag != "" {
		flgs := Flags{}
		err := flgs.Capture([]string{flag})
		if err != nil {
			panic(err)
		}
		rwRule.Flags = &flgs
	}
	rwRule.Params = params

	data, _ := json.Marshal(rwRule)
	return string(data)
}

func TestParserRewriteRule(t *testing.T) {
	//TODO: parser cannot handle `RewriteRule ^search/(.*)$ "[a-zA-Z]"`
	//Might need to be its own handler like RewriteCond
	tests := [][]string{
		{`RewriteRule "^search/(.*)$" "/search.php?term=$1" "[B= ?]"`,
			generateDirectiveString("RewriteRule", "[B= ?]", "^search/(.*)$", "/search.php?term=$1")},
		{`RewriteRule "^search/(.*)$" "/search.php?term=$1" [B=?]`,
			generateDirectiveString("RewriteRule", "[B=?]", "^search/(.*)$", "/search.php?term=$1")},
		{`RewriteRule ^search/(.*)$ [a-zA-Z]+`,
			generateDirectiveString("RewriteRule", "", "^search/(.*)$", "[a-zA-Z]+")},
		{`RewriteRule ^search/(.*)$ "[a-zA-Z]"`,
			generateDirectiveString("RewriteRule", "", "^search/(.*)$", "[a-zA-Z]")},
	}

	testMatch(t, tests)
}

func TestParserSetEnvIf(t *testing.T) {
	tests := [][]string{
		{`SetEnvIf user-agent "WebBandit" stayout=1`,
			""},
	}

	testMatch(t, tests)
}

func TestParserDirectoryIndex(t *testing.T) {
	tests := [][]string{
		{`RewriteEngine On
		  DirectoryIndex index.php index.html`,
			""},
		{`Require all granted
		  Require not env stayout`,
			""},
	}

	testMatch(t, tests)
}

func TestParserRewriteCond(t *testing.T) {

}

func TestParserIfThenElse(t *testing.T) {

}

func TestParserInvalidDirective(t *testing.T) {

}

func TestParserFiles(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../testFiles")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		path := dir + "/" + f.Name()
		_, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := ioutil.ReadFile(path)
		_, err = htaccessParser.ParseBytes(f.Name(), b)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
	}
}
