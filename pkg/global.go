/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package pkg

import "github.com/spf13/cobra"

type CommandOptions interface {
	Exe(parameter []string)
}

type AbstractFactory interface {
	Option() CommandOptions
}

var CobraCat = &cobra.Command{
	Use:   "cat [string to cat]",
	Short: "Create the DIRECTORY(ies), if they do not already exist.",
	Long:  `Create the DIRECTORY(ies), if they do not already exist.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory AbstractFactory
			cmdfactory = new(CatFactory)

			var option CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

var CobraLs = &cobra.Command{
	Use:   "ls [string to ls]",
	Short: "List information about the FILEs (the current directory by default).",
	Long:  `List information about the FILEs (the current directory by default).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory AbstractFactory
			cmdfactory = new(LsFactory)

			var option CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

var CobraTouch = &cobra.Command{
	Use:   "touch [string to touch]",
	Short: "Get content from standard input and write it to file.",
	Long:  `Get content from standard input and write it to file.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory AbstractFactory
			cmdfactory = new(TouchFactory)

			var option CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

var CobraMkdir = &cobra.Command{
	Use:   "mk [string to mk]",
	Short: "Create the DIRECTORY(ies), if they do not already exist.",
	Long:  `Create the DIRECTORY(ies), if they do not already exist.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory AbstractFactory
			cmdfactory = new(MkdirFactory)

			var option CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

var CobraRm = &cobra.Command{
	Use:   "rm [string to rm]",
	Short: "Remove (unlink) the FILE(s).",
	Long:  `Remove (unlink) the FILE(s).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			var cmdfactory AbstractFactory
			cmdfactory = new(RmdirFactory)

			var option CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

var CobraMv = &cobra.Command{
	Use:   "mv [string to mv]",
	Short: "Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.",
	Long:  `Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory AbstractFactory
			cmdfactory = new(MvFactory)

			var option CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}

var CobraTar = &cobra.Command{
	Use:   "tar [string to tar]",
	Short: "The default action is to add or replace tar file entries from list",
	Long:  `The default action is to add or replace tar file entries from list`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			var cmdfactory AbstractFactory
			cmdfactory = new(TarFactory)

			var option CommandOptions
			option = cmdfactory.Option()
			option.Exe(args)
		}
	},
}
