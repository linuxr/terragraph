package configs

type Resource struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"` // 资源类型
}
