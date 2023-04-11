package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var ip string
var port string
var username string
var password string
var database string
var out_dir string
var table string

var rootCmd = &cobra.Command{
	Use:   "export",
	Short: "export是一款sql导出工具",
	Long:  `export是一款由cobra开发的命令行工具，用于导出多种数据库的转储sql文件`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "localhost", "主机名")
	rootCmd.PersistentFlags().StringVar(&port, "port", "3306", "端口")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "用户名")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "root", "密码")
	rootCmd.PersistentFlags().StringVar(&database, "database", "first", "数据库名")
	rootCmd.PersistentFlags().StringVarP(&out_dir, "out_dir", "o", "./data/", "文件输出目录")
}
