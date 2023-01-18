/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package cmd

import (
	"kube-tools/pkg"

	"github.com/spf13/cobra"
)

// touchCmd represents the touch command
var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "Get content from standard input and write it to file.",
	Long:  `Get content from standard input and write it to file.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.TouchFactory)

			var option pkg.CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(touchCmd)
}
