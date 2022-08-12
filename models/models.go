package models

type Node struct {
	Id            string   `json:"id"`            // 唯一标识
	Name          string   `json:"name"`          // 资源名称
	ProviderName  string   `json:"providerName"`  // provider名称
	ModuleAddress string   `json:"moduleAddress"` // 模块address
	Address       string   `json:"address"`       // 资源address
	Type          string   `json:"type"`          // 资源类型，和 provider 相关
	DependsOn     []string `json:"dependsOn"`     // 依赖的资源
	Group         string   `json:"group"`         // 分组
	IsDisplay     bool     `json:"isDisplay"`     // 是否显示
}

type Edge struct {
	SourceId string `json:"sourceId"`
	TargetId string `json:"targetId"`
}
