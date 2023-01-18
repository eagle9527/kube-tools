/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package cmd

import (
	"kube-tools/pkg"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove (unlink) the FILE(s).",
	Long:  `Remove (unlink) the FILE(s).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.RmdirFactory)

			var option pkg.CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
