package parsers

import (
	"github.com/linuxr/terragraph/models"
	"github.com/linuxr/terragraph/utils"
	"github.com/tidwall/gjson"
)

func ParseResources(module gjson.Result, nodes []models.Node) ([]models.Node, error) {
	var err error
	if nodes == nil {
		nodes = make([]models.Node, 0)
	}

	resources := module.Get("resources")
	// fmt.Printf("resources: %v\n", resources)

	if resources.Exists() && resources.IsArray() {
		moduleAddr := module.Get("address").String()

		for _, res := range resources.Array() {
			node := models.Node{
				Id:            utils.GenUUID(),
				Name:          res.Get("name").String(),
				ProviderName:  res.Get("provider_name").String(),
				ModuleAddress: moduleAddr,
				Address:       res.Get("address").String(),
				Type:          res.Get("type").String(),
				DependsOn:     make([]string, 0),
			}

			if res.Get("depends_on").IsArray() {
				for _, depend := range res.Get("depends_on").Array() {
					node.DependsOn = append(node.DependsOn, depend.String())
				}
			}

			nodes = append(nodes, node)
		}
	}

	childModules := module.Get("child_modules")
	if childModules.Exists() && childModules.IsArray() {
		for _, cModule := range childModules.Array() {
			// fmt.Printf("child module: %v\n", cModule)
			nodes, err = ParseResources(cModule, nodes)
			if err != nil {
				return nodes, err
			}
		}
	}

	return nodes, nil
}
