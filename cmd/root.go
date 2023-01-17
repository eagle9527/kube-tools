/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"kube-tools/pkg"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kube-tools",
	Short: "kube-tools",
	Long:  `kube-tools`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var rootCmd = &cobra.Command{Use: "kube-tools"}
	rootCmd.AddCommand(
		pkg.CobraLs,
		pkg.CobraTouch,
		pkg.CobraMv,
		pkg.CobraTar,
		pkg.CobraCat,
		pkg.CobraRm,
		pkg.CobraMkdir,
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
