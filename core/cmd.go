package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// for help
var longDesc = `i18n-message-gen is a tool, for generate i18n messages from golang source file.

Using it, we can define the tags and values in the i18n configuration file as referable golang variables and comments, and automatically extract them into the configuration file.
built with love by humbinal in Go.

Complete documentation is available at https://github.com/humbinal/i18n-message-gen`

// define args
var source string
var target string
var name string
var verbose bool

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "i18n-message-gen",
		Short: "i18n message file generator form golang source file.",
		Long:  longDesc,
		Run: func(cmd *cobra.Command, args []string) {
			pwd, _ := os.Getwd()
			if verbose {
				println("i18n-message-gen input args: ")
				println("  source:", source)
				println("  target:", target)
				println("  name:", name)
				println("  working dir:", pwd)
			}
			// 判斷相对目录还是绝对目录
			if !filepath.IsAbs(source) {
				source = filepath.Join(pwd, source)
				if verbose {
					println("source is relative path, full path is:", source)
				}
			}
			// 文件或目录是否存在
			_, err := os.Stat(source)
			if err != nil {
				println("source file not exist:", source)
				os.Exit(1)
			}
			Generate(source, target, name)
		},
	}
	rootCmd.Flags().StringVarP(&source, "source", "s", ".", "source dir or file for generate.")
	rootCmd.Flags().StringVarP(&target, "target", "t", ".", "output dir for generated i18n file.")
	rootCmd.Flags().StringVarP(&name, "name", "n", "zh-CN.json", "output file name for generated i18n file.")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "output file name for generated i18n file.")
	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
