/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"kube-tools/pkg"
	"os"
)

var cmdLs = &cobra.Command{
	Use:   "ls [string to ls]",
	Short: "List information about the FILEs (the current directory by default).",
	Long:  `List information about the FILEs (the current directory by default).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.LsFactory)

			var option pkg.CommandOptions
			option = cmdfactory.CreateCommandOptions()
			option.Exe(args)
		}
	},
}

var cmdCat = &cobra.Command{
	Use:   "cat [string to cat]",
	Short: "Create the DIRECTORY(ies), if they do not already exist.",
	Long:  `Create the DIRECTORY(ies), if they do not already exist.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.CatFactory)

			var option pkg.CommandOptions
			option = cmdfactory.CreateCommandOptions()
			option.Exe(args)
		}
	},
}

var cmdMkdir = &cobra.Command{
	Use:   "mk [string to mk]",
	Short: "Create the DIRECTORY(ies), if they do not already exist.",
	Long:  `Create the DIRECTORY(ies), if they do not already exist.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.MkdirFactory)

			var option pkg.CommandOptions
			option = cmdfactory.CreateCommandOptions()
			option.Exe(args)
		}
	},
}

var cmdRm = &cobra.Command{
	Use:   "rm [string to rm]",
	Short: "Remove (unlink) the FILE(s).",
	Long:  `Remove (unlink) the FILE(s).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.RmdirFactory)

			var option pkg.CommandOptions
			option = cmdfactory.CreateCommandOptions()
			option.Exe(args)
		}
	},
}

var cmdTouch = &cobra.Command{
	Use:   "touch [string to touch]",
	Short: "Get content from standard input and write it to file.",
	Long:  `Get content from standard input and write it to file.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.TouchFactory)

			var option pkg.CommandOptions
			option = cmdfactory.CreateCommandOptions()
			option.Exe(args)
		}
	},
}

var cmdMv = &cobra.Command{
	Use:   "mv [string to mv]",
	Short: "Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.",
	Long:  `Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.MvFactory)

			var option pkg.CommandOptions
			option = cmdfactory.CreateCommandOptions()
			option.Exe(args)
		}
	},
}

var cmdTar = &cobra.Command{
	Use:   "tar [string to tar]",
	Short: "The default action is to add or replace tar file entries from list",
	Long:  `The default action is to add or replace tar file entries from list`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory pkg.AbstractFactory
			cmdfactory = new(pkg.TarFactory)

			var option pkg.CommandOptions
			option = cmdfactory.CreateCommandOptions()
			option.Exe(args)
		}
	},
}

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
	rootCmd.AddCommand(cmdLs, cmdCat, cmdMkdir, cmdRm, cmdTouch, cmdMv, cmdTar)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
