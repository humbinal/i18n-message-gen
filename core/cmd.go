package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var source string
var target string
var name string

var rootCmd = &cobra.Command{
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

func init() {
	rootCmd.Flags().StringVarP(&source, "source", "s", "examples", "source dir or file for generate.")
	rootCmd.Flags().StringVarP(&target, "target", "t", "examples", "output dir for generated i18n file.")
	rootCmd.Flags().StringVarP(&name, "name", "n", "zh-CN.json", "output file name for generated i18n file.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
