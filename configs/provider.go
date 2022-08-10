package configs

import (
	"bytes"
	"io/ioutil"

	"github.com/spf13/viper"
)

type ProviderType string

const (
	ProviderAlicloud ProviderType = "alicloud"
)

type Providers struct {
	Providers []Provider `mapstructure:"providers"`
}

type Provider struct {
	Type      string        `mapstructure:"type"`
	Resources []ProResource `mapstructure:"resources"`
}

type ProResource struct {
	Name    string `mapstructure:"name"`
	Address string `mapstructure:"address"`

	IsDisplay bool   `mapstructure:"isDisplay"` // 资源类型是否显示
	Group     string `mapstructure:"group"`     // 资源类型分组的名称
	Type      string `mapstructure:"type"`      // 资源类型
}

// GetDefaultProvider 解析 provider 的配置
func GetDefaultProvider() (*Providers, error) {
	content := MustAsset("configs/default_providers.yaml")
	return getProviders(content)
}

// GetCustomizeProviders 解析自定义配置中的所有 provider
func GetCustomizeProviders(fp string) (*Providers, error) {
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	return getProviders(content)
}

func getProviders(content []byte) (*Providers, error) {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
		return nil, err
	}

	var conf Providers
	err := viper.Unmarshal(&conf)
	return &conf, err

}

// MergeProviders 合并默认的provider配置和自定义的provider配置
func MergeProviders(defaultProviders, providers []Provider) map[string][]ProResource {
	mResources := make(map[string][]ProResource)

	// 默认资源配置
	for _, p := range defaultProviders {
		mResources[p.Type] = p.Resources
	}

	// 自定义配置覆盖默认配置
	for _, p := range providers {
		if _, ok := mResources[p.Type]; !ok {
			mResources[p.Type] = p.Resources
			continue
		}

		mResources[p.Type] = MergeResources(mResources[p.Type], p.Resources)
	}

	return mResources
}

// MergeResources 合并默认的资源配置和自定义的资源配置
func MergeResources(defaultResources, resources []ProResource) []ProResource {
	var results = make([]ProResource, 0)
	for _, resX := range defaultResources {
		isFind := false
		for _, resY := range resources {
			if resX.Address == resY.Address {
				isFind = true
				break
			}
		}

		if !isFind {
			results = append(results, resX)
		}
	}

	results = append(results, resources...)
	return results
}
