package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

const DefJSONName = "CloudFormationResourceSpecification.json"

var DefJSONMap = map[string]string{
	"ap-northeast-1": "https://doigdx0kgq9el.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
}

type CfnDef struct {
	PropertyTypes                map[string]PropertyType `json:"PropertyTypes"`
	ResourceTypes                map[string]ReourceType  `json:"ResourceTypes"`
	ResourceSpecificationVersion string                  `json:"ResourceSpecificationVersion"`
}

type PropertyType struct {
	Documentation string              `json:"Documentation"`
	Properties    map[string]Property `json:"Properties"`
}

type Property struct {
	Documentation     string `json:"Documentation"`
	DuplicatesAllowed bool   `json:"DuplicatesAllowed"`
	ItemType          string `json:"ItemType"`
	PrimitiveItemType string `json:"PrimitiveItemType"`
	PrimitiveType     string `json:"PrimitiveType"`
	Required          bool   `json:"Required"`
	Type              string `json:"Type"`
	UpdateType        string `json:"UpdateType"`
}

type ReourceType struct {
	Documentation string               `json:"Documentation"`
	Attributes    map[string]Attribute `json:"Attributes"`
	Properties    map[string]Property  `json:"Properties"`
}

type Attribute struct {
	ItemType          string `json:"ItemType"`
	PrimitiveItemType string `json:"PrimitiveItemType"`
	PrimitiveType     string `json:"PrimitiveType"`
	Type              string `json:"Type"`
}

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
			Action:  update,
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
			Action:  parse,
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

func parse(ctx *cli.Context) {
	region := ctx.String("region")
	fpath := filepath.Join(region, DefJSONName)
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err.Error())
	}

	cfnDef := CfnDef{}
	err = json.Unmarshal(bytes, &cfnDef)
	if err != nil {
		fmt.Printf("ERROR:%v", err.Error())
	}
	// DEBUG
	//for ptName, pt := range cfnDef.PropertyTypes {
	//	fmt.Printf("%s", ptName)
	//	fmt.Printf("%s", pt.Documentation)
	//	for pName, p := range pt.Properties {
	//		fmt.Println(pName)
	//		fmt.Println(p)
	//	}
	//}
}

func update(ctx *cli.Context) {
	region := ctx.String("region")
	url := DefJSONMap[region]
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
