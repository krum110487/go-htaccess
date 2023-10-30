package htaccess

import (
	"github.com/krum110487/go-htaccess/config"
	"github.com/krum110487/go-htaccess/parser"
)

type Context struct {
	requestConfig config.CoreConfig
	runtimeStack  []parser.DirectiveEntry
	stackGroup    map[string][]parser.DirectiveEntry
}

func (c *Context) GetGroup(name string) []parser.DirectiveEntry {
	//Do stuff here
	return []parser.DirectiveEntry{}
}

func (c *Context) AddToGroup(name string) {
	//Add item to map with the results of processing it.
}

func (c *Context) ClearGroup(name string) {
	//Empty the array
}

func (c *Context) DeleteGroup(name string) {
	//Delete the map item.
}

func (c *Context) EvaluateGroupBool(name string, bottomUp bool) bool {
	//Loops through the array and determines if the Group passes.
	//Check each BooleanResult from the ProcessResults, Or'd or And'd against the next result.
	//Bottom up simply means traverse the stack backwards.
	return false
}
