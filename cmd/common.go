package cmd

import (
	"github.com/urfave/cli"
)

const (
	DefDir          = "def"
	SnippetsDir     = "snippets"
	ResourceSnippet = `
snippet %s
Type %s
Properties%s
endsnippet
`
	PropertySnippet = `
# %s
snippet %s
Type %s
Properties%s
endsnippet
`
)

var Commands = []cli.Command{
	cmdUpdate,
	cmdParse,
	cmdCreate,
}
