package models

type Node struct {
	Id            string   // 唯一标识
	Name          string   // 资源名称
	ProviderName  string   // provider名称
	ModuleAddress string   // 模块address
	Address       string   // 资源address
	Type          string   // 资源类型，和 provider 相关
	DependsOn     []string // 依赖的资源
	ComboId       string   // 分组的ID
}

type Edge struct {
	SourceId string
	TargetId string
}
