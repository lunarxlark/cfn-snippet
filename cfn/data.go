package cfn

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
