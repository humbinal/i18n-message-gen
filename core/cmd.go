package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var source string
var target string
var name string

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "i18n-message-gen",
		Short: "i18n message generator for golang",
		Long: `i18n-message-gen 
    Complete documentation is available at https://github.com/humbinal/i18n-message-gen`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("source:", source)
			fmt.Println("target:", target)
			fmt.Println("name:", name)
			Gen(source, target+"/"+name)
		},
	}
	rootCmd.Flags().StringVarP(&source, "source", "s", ".", "source dir or file for generate, default current dir.")
	rootCmd.Flags().StringVarP(&target, "target", "t", ".", "output dir for generated i18n file, default current dir.")
	rootCmd.Flags().StringVarP(&name, "name", "n", "zh-CN.json", "output file name for generated i18n file.")
	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
