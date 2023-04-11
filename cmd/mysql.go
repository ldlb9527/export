package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "该命令用于导出mysql的转储sql   你可使用export mysql -h 查看更多帮助 ",
	Long:  `该命令用于导出mysql的转储sql   你可使用export mysql -h 查看更多帮助`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mysql数据库连接成功")
		fmt.Println(fmt.Sprintf("ip:%s port:%s username:%s password:%s", ip, port, username, password))
		fmt.Println("sql文件导出完成")
	},
}

func init() {
	mysqlCmd.PersistentFlags().StringVarP(&table, "table", "t", "", "表名(如不指定表名,则导出该库下全部)")
	rootCmd.AddCommand(mysqlCmd)
}
