package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	inputPath  string
	provider   string
	configPath string

	rootCmd = &cobra.Command{
		Use: "terragraph",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("input: %s\n", inputPath)
			fmt.Printf("provider: %s\n", provider)
			fmt.Printf("config: %s\n", configPath)

			// TODO: 解析nodes信息
			// TODO: 解析配置信息
			// TODO: 将配置信息覆盖nodes的配置
			// TODO: 根据nodes生成edges

			// TODO: 导出json用于G6绘图
			return nil
		},
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&inputPath, "input", "i", "", "输入文件的路径(HCL, tfplan.json, tfstate.json)")

	rootCmd.Flags().StringVar(&provider, "provider", "alicloud", "provider 类型")
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "自定义配置文件的路径")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("config")
}
