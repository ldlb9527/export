package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "打印版本信息",
	Long:  `打印export自身的版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
