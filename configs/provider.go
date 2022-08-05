package configs

import (
	"bytes"

	"github.com/linuxr/terragraph/configs/alicloud"
	"github.com/spf13/viper"
)

type ProviderType string

const (
	ProviderAlicloud ProviderType = "alicloud"
)

type Provider struct {
	Type      string        `mapstructure:"type"`
	Resources []ProResource `mapstructure:"resources"`
}

type ProResource struct {
	Name string `mapstructure:"name"`

	IsDisplay bool   `mapstructure:"isDisplay"` // 资源类型是否显示
	Group     string `mapstructure:"group"`     // 资源类型分组的名称
	Type      string `mapstructure:"type"`      // 资源类型
}

func providerSettings(providerType ProviderType) []byte {
	if providerType == ProviderAlicloud {
		return alicloud.MustAsset("configs/alicloud/provider.yaml")
	}

	return nil
}

// GetDefaultProvider 解析 provider 的配置
func GetDefaultProvider(providerType ProviderType) (*Provider, error) {
	content := providerSettings(providerType)
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
		return nil, err
	}

	var conf Provider
	err := viper.Unmarshal(&conf)
	return &conf, err
}
