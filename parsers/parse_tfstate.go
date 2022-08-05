package parsers

import (
	"github.com/linuxr/terragraph/models"
	"github.com/linuxr/terragraph/utils"
	"github.com/tidwall/gjson"
)

func ParseNodesFromTfState(module gjson.Result, nodes []models.Node) ([]models.Node, error) {
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
			nodes, err = ParseNodesFromTfState(cModule, nodes)
			if err != nil {
				return nodes, err
			}
		}
	}

	return nodes, nil
}

func ParseEdgeFromNodes(nodes []models.Node) ([]models.Edge, error) {
	edges := make([]models.Edge, 0)
	// 创建 nodes map，address 作为key，用于判断 depends_on 依赖的是否是资源
	m := make(map[string]string)
	for _, n := range nodes {
		m[n.Address] = n.Id
	}

	for _, n := range nodes {
		for _, depend := range n.DependsOn {
			if id, ok := m[depend]; ok {
				edges = append(edges, models.Edge{
					SourceId: n.Id,
					TargetId: id,
				})
			}
		}
	}

	return edges, nil
}
