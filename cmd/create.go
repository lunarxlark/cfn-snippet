package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lunarxlark/cfn-snippet/cfn"
	"github.com/urfave/cli"
)

var cmdCreate = cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "create snippet from CloudFormation definition json",
	Action:  doCreate,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "region", Usage: "which region's cfn definition do you get.", Value: "ap-northeast-1"},
	},
}

func doCreate(ctx *cli.Context) error {
	region := ctx.String("region")
	snippetsRegionDir := filepath.Join(SnippetsDir, region)
	defRegionDir := filepath.Join(DefDir, region)
	if err := os.MkdirAll(snippetsRegionDir, 0755); err != nil {
		return fmt.Errorf("%w", err)
	}

	f, err := os.Create(filepath.Join(snippetsRegionDir, "cfn.snippets"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer f.Close()

	defFilePath := filepath.Join(defRegionDir, "CloudFormationResourceSpecification.json")
	bytes, err := ioutil.ReadFile(defFilePath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	cfnDef := cfn.CfnDef{}
	if err = json.Unmarshal(bytes, &cfnDef); err != nil {
		return fmt.Errorf("%w", err)
	}

	// Resource
	for resourceTypeName, resourceType := range cfnDef.ResourceTypes {
		pOutput := ""
		i := 0
		for propertyName, property := range resourceType.Properties {
			if property.Type != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.Type)
			} else if property.ItemType != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.ItemType)
			} else if property.PrimitiveType != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.PrimitiveType)
			} else if property.PrimitiveItemType != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.PrimitiveItemType)
			}
			i++
		}
		_, err = io.WriteString(f, fmt.Sprintf(ResourceSnippet, resourceTypeName, resourceTypeName, pOutput))
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	// Property
	for propertyTypeName, propertyType := range cfnDef.PropertyTypes {

		pOutput := ""
		i := 0
		for propertyName, property := range propertyType.Properties {
			if property.Type != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.Type)
			} else if property.ItemType != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.ItemType)
			} else if property.PrimitiveType != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.PrimitiveType)
			} else if property.PrimitiveItemType != "" {
				pOutput = fmt.Sprintf("%s\n\t%s: ${%d:%s}", pOutput, propertyName, i, property.PrimitiveItemType)
			}
			i++
		}
		_, err = io.WriteString(f, fmt.Sprintf(PropertySnippet, propertyType.Documentation, propertyTypeName, propertyTypeName, pOutput))
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	fmt.Printf("success to create snippets of CloudFormation in %s", region)
	return nil
}
