package models

type Provider struct {
	Name      string
	Resources []ProResource
}

type ProResource struct {
	Name string

	IsDisplay bool   // 资源是否显示
	Group     string // 资源分组的名称
	Type      string // 资源类型
}

type Resource struct {
	Name string
}
