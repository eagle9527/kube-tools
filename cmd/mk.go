/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package cmd

import (
	"kube-tools/pkg"

	"github.com/spf13/cobra"
)

// mkCmd represents the mk command
var mkCmd = &cobra.Command{
	Use:   "mk",
	Short: "Create the DIRECTORY(ies), if they do not already exist.",
	Long:  `Create the DIRECTORY(ies), if they do not already exist.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.MkdirFactory)

			var option pkg.CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(mkCmd)
}
