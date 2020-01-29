package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lunarxlark/cfn-snippet/cfn"
	"github.com/urfave/cli"
)

var cmdUpdate = cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Usage:   "update cloudformation definition by getting from link in aws cfn doc.",
	Action:  doUpdate,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "region", Usage: "which region's cfn definition do you get.", Value: "ap-northeast-1"},
	},
}

func doUpdate(ctx *cli.Context) error {
	region := ctx.String("region")
	url := cfn.DefJSONMap[region]
	defJSONName := filepath.Base(url)
	defRegionDir := filepath.Join(DefDir, region)
	if err := os.MkdirAll(defRegionDir, 0755); err != nil {
		fmt.Errorf("%w", err)
	}

	f, err := os.Create(filepath.Join(defRegionDir, defJSONName))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer f.Close()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	_, err = io.WriteString(f, string(b))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Printf("success to get definition JSON of CloudFormation in %s", region)
	return nil
}
