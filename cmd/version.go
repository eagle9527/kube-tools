package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kube-tools.",
	Long:  `Print the version number of kube-tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kube-tools Version: v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
