package cmds

import (
	"encoding/json"
	"io/ioutil"

	"github.com/linuxr/terragraph/configs"
	"github.com/linuxr/terragraph/models"
	"github.com/linuxr/terragraph/parsers"
	"github.com/spf13/cobra"
)

var (
	inputPath  string
	configPath string

	rootCmd = &cobra.Command{
		Use: "terragraph",
		RunE: func(cmd *cobra.Command, args []string) error {
			// fmt.Printf("input: %s\n", inputPath)
			// fmt.Printf("config: %s\n", configPath)

			// 解析nodes信息
			nodes, err := parsers.ParseTfState(inputPath)
			if err != nil {
				return err
			}

			// 解析配置信息
			defaultProviders, err := configs.GetDefaultProvider()
			if err != nil {
				return err
			}
			providers, err := configs.GetCustomizeProviders(configPath)
			if err != nil {
				return err
			}

			mergedProviders := configs.MergeProviders(defaultProviders.Providers, providers.Providers)

			// fmt.Printf("mergedProviders: %v\n", mergedProviders)

			// 将配置信息覆盖nodes的配置
			nodes, err = parsers.ParseNodesWithConfig(nodes, mergedProviders)
			if err != nil {
				return err
			}
			// fmt.Printf("count: %d\nnodes: %v\n", len(nodes), nodes)

			// 根据nodes生成edges
			edges, err := parsers.ParseEdgeFromNodes(nodes)
			if err != nil {
				return err
			}
			// fmt.Printf("edges: %v\n", edges)

			// 导出json用于G6绘图
			var data = struct {
				Nodes  []models.Node `json:"nodes"`
				Edges  []models.Edge `json:"edges"`
				Groups []string      `json:"groups"`
			}{
				Nodes:  nodes,
				Edges:  edges,
				Groups: parsers.ParseGroupsFromNodes(nodes),
			}

			content, err := json.Marshal(&data)
			if err != nil {
				return err
			}

			return ioutil.WriteFile("./data.json", content, 0644)
		},
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&inputPath, "input", "i", "", "输入文件的路径(HCL, tfplan.json, tfstate.json)")

	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "自定义配置文件的路径")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("config")
}
