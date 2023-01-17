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
		cmdfactory := new(pkg.CmdFactory)
		cmdOption := cmdfactory.CreateCommandOption("ls")
		if len(args) == 1 {
			cmdOption.Exe(args)
		}
	},
}

var cmdCat = &cobra.Command{
	Use:   "cat [string to cat]",
	Short: "Create the DIRECTORY(ies), if they do not already exist.",
	Long:  `Create the DIRECTORY(ies), if they do not already exist.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdfactory := new(pkg.CmdFactory)
		cmdOption := cmdfactory.CreateCommandOption("cat")
		if len(args) == 1 {
			cmdOption.Exe(args)
		}
	},
}

var cmdMkdir = &cobra.Command{
	Use:   "mk [string to mk]",
	Short: "Create the DIRECTORY(ies), if they do not already exist.",
	Long:  `Create the DIRECTORY(ies), if they do not already exist.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdfactory := new(pkg.CmdFactory)
		cmdOption := cmdfactory.CreateCommandOption("mkdir")
		if len(args) == 1 {
			cmdOption.Exe(args)
		}
	},
}

var cmdRm = &cobra.Command{
	Use:   "rm [string to rm]",
	Short: "Remove (unlink) the FILE(s).",
	Long:  `Remove (unlink) the FILE(s).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdfactory := new(pkg.CmdFactory)
		cmdOption := cmdfactory.CreateCommandOption("rm")
		if len(args) == 1 {
			cmdOption.Exe(args)
		}
	},
}

var cmdTouch = &cobra.Command{
	Use:   "touch [string to touch]",
	Short: "Get content from standard input and write it to file.",
	Long:  `Get content from standard input and write it to file.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cmdfactory := new(pkg.CmdFactory)
		cmdOption := cmdfactory.CreateCommandOption("touch")
		if len(args) == 2 {
			cmdOption.Exe(args)
		}
	},
}

var cmdMv = &cobra.Command{
	Use:   "mv [string to mv]",
	Short: "Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.",
	Long:  `Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cmdfactory := new(pkg.CmdFactory)
		cmdOption := cmdfactory.CreateCommandOption("mv")
		if len(args) == 2 {
			cmdOption.Exe(args)
		}
	},
}

var cmdTar = &cobra.Command{
	Use:   "tar [string to tar]",
	Short: "The default action is to add or replace tar file entries from list",
	Long:  `The default action is to add or replace tar file entries from list`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cmdfactory := new(pkg.CmdFactory)
		cmdOption := cmdfactory.CreateCommandOption("tar")
		if len(args) == 2 {
			cmdOption.Exe(args)
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

func init() {
}
