/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package cmd

import (
	"kube-tools/pkg"

	"github.com/spf13/cobra"
)

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.",
	Long:  `Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.MvFactory)

			var option pkg.CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)
}
