package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lunarxlark/cfn-snippet/cfn"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Usage = "test usage"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update cloudformation definition by getting from link in aws cfn doc.",
			Action:  cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "region",
					Usage: "which region's cfn definition do you get.",
					Value: "ap-northeast-1",
				},
			},
		},
		{
			Name:    "parse",
			Aliases: []string{"p"},
			Usage:   "parse CFn def json",
			Action:  cmdParse,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "region",
					Usage: "which region's cfn definition do you get.",
					Value: "ap-northeast-1",
				},
			},
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create snippet from CloudFormation definition json",
			Action:  cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "region",
					Usage: "which region's cfn definition do you get.",
					Value: "ap-northeast-1",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("fatal")
	}
}

func cmdCreate(ctx *cli.Context) {
	//region := ctx.String("region")
}

func cmdParse(ctx *cli.Context) {
	region := ctx.String("region")
	fpath := filepath.Join(region, cfn.DefJSONName)
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err.Error())
	}

	cfnDef := cfn.CfnDef{}
	err = json.Unmarshal(bytes, &cfnDef)
	if err != nil {
		fmt.Printf("ERROR:%v", err.Error())
	}
}

func cmdUpdate(ctx *cli.Context) {
	region := ctx.String("region")
	url := cfn.DefJSONMap[region]
	defJSONName := filepath.Base(url)
	err := os.Mkdir(region, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}

	f, err := os.Create(filepath.Join(region, defJSONName))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = io.WriteString(f, string(b))
	if err != nil {
		fmt.Println(err.Error())
	}
}
