package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var oracleCmd = &cobra.Command{
	Use:   "oracle",
	Short: "该命令用于导出oracle的转储sql   你可使用export oracle -h 查看更多帮助 ",
	Long:  `该命令用于导出oracle的转储sql   你可使用export oracle -h 查看更多帮助`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oracle数据库连接成功")
		fmt.Println(fmt.Sprintf("ip:%s port:%d username:%s password:%s", ip, port, username, password))
		fmt.Println("sql文件导出完成")
	},
}

func init() {
	rootCmd.AddCommand(oracleCmd)
}
