package parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var htaLexer = lexer.MustStateful(lexer.Rules{
	"Root": {
		{"LineContinuation", `\/[\r\n]`, nil},
		{"EOL", `[\r\n]`, nil},
		{"Comment", `#[^\r\n]*`, nil},
		//{"Flag", `\[[^\r\n\s\]]*\]\s*$`, nil},
		//{"QFlag", `"\[[^\r\n\]]*\]\s*"\s*$`, nil},
		{"QStr", `"(\\"|[^\r\n"])*"`, nil},
		{"Whitespace", `[ \t]+`, nil},
		{"TagClose", `\<\/`, lexer.Push("TagClose")},
		{"TagOpen", `\<`, lexer.Push("TagOpen")},
		{"DirFlag", directivesWithFlagRegex, lexer.Push("DirFlag")},
		{"Dir", directiveRegex, nil},
		{"Str", `[^\s\n\r]+`, nil},
	},
	"DirFlag": {
		{"LineContinuation", `\/[\r\n]`, nil},
		{"EOL", `[\r\n]`, lexer.Pop()},
		{"QEqFlag", `("(\\"|[^\r\n"])+=(\\"|[^\r\n"])+"|'(\\'|[^\r\n'])+=(\\'|[^\r\n'])+')`, nil},
		{"QStr", `"(\\"|[^\r\n"])*"`, nil},
		{"Whitespace", `[ \t]+`, nil},
		{"EqFlag", `[^\r\n \t]+=[^\r\n \t]+`, nil},
		{"Str", `[^\s\n\r]+`, nil},
	},
	"TagOpen": {
		{"LineContinuation", `\/[\r\n]`, nil},
		{"QStr", `"(\\"|[^\r\n"])*"`, nil},
		{"TagEnd", `>`, lexer.Pop()},
		{"Whitespace", `[ \t]+`, nil},
		{"TagDir", tagDirectiveRegex, nil},
		{"TagStr", `[^\s\n\r<>]+`, nil},
	},
	"TagClose": {
		{"TagEnd", `>`, lexer.Pop()},
		{"TagCloseName", `[A-Za-z0-9]+`, nil},
		{"Whitespace", `[ \t]+`, nil},
	},
})

type HtaccessAST struct {
	Dirs []*DirectiveEntry `parser:"@@*" json:"directives"`
}

type DirectiveEntry struct {
	Pos         lexer.Position `parser:""                                                json:"pos,omitempty"`
	IfGroup     *IfEntry       `parser:"(   @@"                                          json:"ifgroup,omitempty"`
	Tag         *TagEntry      `parser:"  | @@"                                          json:"tag,omitempty"`
	SetEnvIf    *SetEnvIf      `parser:"  | @@"                                          json:"setEnvIf,omitempty"`
	RewriteCond *RWCondEntry   `parser:"  | @@"                                          json:"rewriteCond,omitempty"`
	Name        string         `parser:"  | @Dir"                                        json:"name,omitempty"`
	Params      []string       `parser:"    @(QStr|Str) @(QStr|Str|QFlag|Flag)?" json:"params,omitempty"`
	Flags       *Flags         `parser:"    @(QFlag|Flag)?    )"                         json:"flags,omitempty"`
	Results     bool           `parser:""                                                json:"results,omitempty"`
}

type RWCondEntry struct {
	Name       string       `parser:"@('RewriteCond')"       json:"name,omitempty"`
	TestString string       `parser:"@(QStr|Str)"            json:"testString,omitempty"`
	Pattern    *CondPattern `parser:"@(QStr|Str|QFlag|Flag)" json:"pattern,omitempty"`
	Flags      *Flags       `parser:"@(QFlag|Flag)*"         json:"flags,omitempty"`
}

type SetEnvIf struct {
	Name     string `parser:"@('SetEnvIf')"   json:"name,omitempty"`
	Variable string `parser:"@(QStr|Str|Dir)" json:"variable,omitempty"`
	Value    string `parser:"@(QStr|Str|Dir)" json:"value,omitempty"`
	Action   string `parser:"@(QStr|Str|Dir)" json:"action,omitempty"`
}

type IfEntry struct {
	Name       string            `parser:"'<' @('If')"        json:"name,omitempty"`
	Expression Expression        `parser:"  @QStr"            json:"expression,omitempty"`
	Dirs       []*DirectiveEntry `parser:"  (   @@"           json:"dirs,omitempty"`
	IfGroup    *IfEntry          `parser:"    | @@"           json:"ifgroup,omitempty"`
	Tags       []*TagEntry       `parser:"    | @@  )*"       json:"tags,omitempty"`
	TagClose   string            `parser:"'</' @TagCloseName" json:"tagclose,omitempty"`
	ElseIf     []*ElseIfEntry    `parser:"(   @@+"            json:"elseifgroup,omitempty"`
	Else       *ElseEntry        `parser:"  | @@  )*"         json:"else,omitempty"`
}

type ElseIfEntry struct {
	Name       string            `parser:"'<' @('ElseIf')"    json:"name,omitempty"`
	Expression Expression        `parser:"  @QStr"            json:"expression,omitempty"`
	Dirs       []*DirectiveEntry `parser:" (   @@"            json:"dirs,omitempty"`
	IfGroup    *IfEntry          `parser:"   | @@"            json:"ifgroup,omitempty"`
	Tags       []*TagEntry       `parser:"   | @@  )*"        json:"tags,omitempty"`
	TagClose   string            `parser:"'</' @TagCloseName" json:"tagclose,omitempty"`
}

type ElseEntry struct {
	Name     string            `parser:"'<' @('Else')"               json:"name,omitempty"`
	Dirs     []*DirectiveEntry `parser:" (   @@"                     json:"dirs,omitempty"`
	IfGroup  *IfEntry          `parser:"   | @@"                     json:"ifgroup,omitempty"`
	Tags     []*TagEntry       `parser:"   | @@  )*"                 json:"tags,omitempty"`
	TagClose string            `parser:"'</' @TagCloseName"          json:"tagclose,omitempty"`
}

type TagEntry struct {
	Name     string            `parser:"'<' @TagDir"        json:"expression,omitempty"`
	Params   []string          `parser:"  @(QStr|TagStr)*"  json:"params,omitempty"`
	Dirs     []*DirectiveEntry `parser:" (   @@"            json:"dirs,omitempty"`
	IfGroup  *IfEntry          `parser:"   | @@"            json:"ifgroup,omitempty"`
	Tags     []*TagEntry       `parser:"   | @@  )*"        json:"tags,omitempty"`
	TagClose string            `parser:"'</' @TagCloseName" json:"tagclose,omitempty"`
}

type Flags struct {
	FlagMap map[string]FlagParams `json:"flagMap,omitempty"`
}

type FlagParams struct {
	NotOp     bool              `json:"notOp,omitempty"`
	ValueStr  string            `json:"valueStr,omitempty"`
	ValueInt  int               `json:"valueInt,omitempty"`
	ValueBool bool              `json:"valueBool,omitempty"`
	ParamMap  map[string]string `json:"paramMap,omitempty"`
}

func (fp *Flags) Capture(values []string) error {
	flagStr := strings.TrimSpace(values[0])
	flagStr = flagStr[1 : len(flagStr)-1]
	flagList := strings.Split(flagStr, ",")
	for _, flag := range flagList {
		params := FlagParams{}
		vals := strings.Split(flag, "=")

		flag := flagNormalized(vals[0])
		if len(vals) > 1 {
			paramList := strings.Split(vals[1], ":")
			switch flag {
			case "CO":
				params = flagMapCookieParams(paramList)
			case "E":
				params = flagMapEnvParams(paramList)
			case "S":
				params = flagConvertToInt(vals[1])
			default:
				params.ValueStr = vals[1]
			}
		} else {
			params.ValueBool = true
		}

		flagMap := make(map[string]FlagParams)
		flagMap[flag] = params
		fp.FlagMap = flagMap
	}

	return nil
}

func (f *Flags) GetFlag(name string) (FlagParams, bool) {
	v, exists := f.FlagMap[strings.ToUpper(name)]
	return v, exists
}

func (f *Flags) Exists(name string) bool {
	_, exists := f.FlagMap[strings.ToUpper(name)]
	return exists
}

func flagNormalized(name string) string {
	abbreviations := map[string]string{
		"BACKREFNOPLUS": "BNP",
		"CHAIN":         "C",
		"COOKIE":        "CO",
		"DISCARDPATH":   "DPI",
		"ENV":           "E",
		"FORBIDDEN":     "F",
		"GONE":          "G",
		"HANDLER":       "H",
		"LAST":          "L",
		"NEXT":          "N",
		"NOCASE":        "NC",
		"NOESCAPE":      "NE",
		"NOSUBREQ":      "NS",
		"PROXY":         "P",
		"PASSTHROUGH":   "PT",
		"QSAPPEND":      "QSA",
		"QSDISCARD":     "QSD",
		"QSLAST":        "QSL",
		"REDIRECT":      "R",
		"SKIP":          "S",
		"TYPE":          "T",
	}
	abbreviation, found := abbreviations[strings.ToUpper(name)]
	if found {
		return abbreviation
	} else {
		return name
	}
}

func flagConvertToInt(value string) FlagParams {
	flagParams := FlagParams{}
	i, err := strconv.Atoi(value)
	if err != nil {
		flagParams.ValueStr = value
	} else {
		flagParams.ValueInt = i
	}
	return flagParams
}

func flagMapEnvParams(paramArr []string) FlagParams {
	params := make(map[string]string)
	flagParams := FlagParams{}
	if len(paramArr) < 2 {
		return flagParams
	}

	varName := paramArr[0]
	if strings.HasPrefix(varName, "!") {
		flagParams.NotOp = true
		varName = strings.TrimPrefix(varName, "!")
	}

	params[varName] = paramArr[1]
	flagParams.ParamMap = params
	return flagParams
}

func flagMapCookieParams(paramArr []string) FlagParams {
	params := make(map[string]string)
	flagParams := FlagParams{}
	if len(paramArr) < 3 {
		return flagParams
	}

	for i, val := range paramArr {
		switch i {
		case 0:
			flagParams.ValueStr = val
		case 1:
			params["domain"] = val
		case 2:
			params["lifetime"] = val
		case 3:
			params["path"] = val
		case 4:
			params["secure"] = val
		case 5:
			params["httponly"] = val
		case 6:
			params["samesite"] = val
		}
	}
	flagParams.ParamMap = params
	return flagParams
}

type CondPattern struct {
	NotOp         bool   `json:"notOp,omitempty"`
	LeftArgument  string `json:"leftArgument,omitempty"`
	Operator      string `json:"operator,omitempty"`
	RightArgument string `json:"rightArgument,omitempty"`
	Results       bool   `json:"results,omitempty"`
}

func (cp *CondPattern) Capture(values []string) error {
	cond := strings.TrimSpace(values[0])
	params := []string{}

	//Step 1 - Check for not op.
	if strings.HasPrefix(cond, "!") {
		cp.NotOp = true
		cond = strings.TrimSpace(strings.TrimPrefix(cond, "!"))
	}

	//Step 1 - Check if it starts with an =,<,>,<=,>=
	opRe := regexp.MustCompile("^(<=|>=|<|>|=)")
	op := opRe.FindString(cond)
	if op != "" {
		params = strings.SplitAfter(cond, op)
		if op == "=" {
			op = "=="
		}
		cp.Operator = op
		cp.RightArgument = params[1]
		return nil
	}

	//Step 2 - Check if it is one of the possible operator, but ONLY operator
	expOpRe := regexp.MustCompile(`^!*\s*-(eq|ge|gt|le|lt|ne|d|f|F|h|l|L|s|U|x)$`)
	expOp := strings.TrimSpace(expOpRe.FindString(cond))
	if expOp != "" {
		cp.Operator = expOp
		return nil
	}

	//Step 4 - Parse the expression.
	re := regexp.MustCompile("[ \t]+")
	parts := re.Split(cond, -1)

	if len(parts) > 3 {
		return errors.New("CondPattern must have at most 3 parts")
	}

	if len(parts) == 1 {
		//Regex test
		cp.Operator = "=~"
		cp.RightArgument = parts[0]
	} else if len(parts) >= 3 {
		//Test with operator
		cp.LeftArgument = parts[0]
		cp.Operator = parts[1]
		cp.RightArgument = parts[2]
	} else {
		//Test without left operator.
		cp.Operator = parts[0]
		cp.RightArgument = parts[1]
	}

	return nil
}

type Expression struct {
	LeftArgument  string `json:"leftArgument,omitempty"`
	Operator      string `json:"operator,omitempty"`
	RightArgument string `json:"rightArgument,omitempty"`
	Results       bool   `json:"results,omitempty"`
}

func (e *Expression) Capture(values []string) error {
	re := regexp.MustCompile("[ \t]+")
	parts := re.Split(values[0], -1)

	if len(parts) > 3 {
		return errors.New("Expression must have at most 3 parts")
	}

	if len(parts) < 2 {
		return errors.New("Expression must have at least 2 parts")
	}

	if len(parts) >= 3 {
		e.LeftArgument = parts[0]
		e.Operator = parts[1]
		e.RightArgument = parts[2]
	} else {
		e.Operator = parts[0]
		e.RightArgument = parts[1]
	}

	return nil
}

func (e *Expression) Resolve() (bool, error) {
	e.Results = false

	switch e.Operator {
	case "==":
		e.Results = e.LeftArgument == e.RightArgument
	case "!=":
		e.Results = e.LeftArgument != e.RightArgument
	case "<=":
		e.Results = e.LeftArgument <= e.RightArgument
	case "<":
		e.Results = e.LeftArgument < e.RightArgument
	case ">":
		e.Results = e.LeftArgument > e.RightArgument
	case ">=":
		e.Results = e.LeftArgument >= e.RightArgument
	case "-z":
		e.Results = e.RightArgument != ""
	case "=~":
		rArgRegex := regexp.MustCompile(e.RightArgument)
		e.Results = rArgRegex.MatchString(e.LeftArgument)
	}

	return e.Results, nil
}

var htaccessParser = participle.MustBuild[HtaccessAST](
	participle.Lexer(htaLexer),
	participle.Elide("LineContinuation", "Comment", "Whitespace"),
	participle.Elide("TagEnd", "TagOpen", "TagClose", "EOL"),
	participle.CaseInsensitive("Dir", "TagDir"),
	trim("Dir"),
	removeQuotes("QStr", "QFlag"),
)

func removeQuotes(types ...string) participle.Option {
	if len(types) == 0 {
		types = []string{"String"}
	}
	return participle.Map(func(t lexer.Token) (lexer.Token, error) {
		value := t.Value
		if len(value) >= 2 {
			value = value[1 : len(value)-1]
		}
		t.Value = value
		return t, nil
	}, types...)
}

func trim(types ...string) participle.Option {
	if len(types) == 0 {
		types = []string{"String"}
	}
	return participle.Map(func(t lexer.Token) (lexer.Token, error) {
		value := strings.TrimSpace(t.Value)
		t.Value = value
		return t, nil
	}, types...)
}

func Parse(code string) (*HtaccessAST, error) {
	return htaccessParser.ParseString("", code)
}
