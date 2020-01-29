package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/lunarxlark/cfn-snippet/cfn"
	"github.com/urfave/cli"
)

var cmdParse = cli.Command{
	Name:    "parse",
	Aliases: []string{"p"},
	Usage:   "parse CFn def json",
	Action:  doParse,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "region", Usage: "which region's cfn definition do you get.", Value: "ap-northeast-1"},
	},
}

func doParse(ctx *cli.Context) error {
	region := ctx.String("region")
	fpath := filepath.Join(region, cfn.DefJSONName)
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	cfnDef := cfn.CfnDef{}
	if err = json.Unmarshal(bytes, &cfnDef); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
