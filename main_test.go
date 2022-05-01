package main

import (
	"github.com/humbinal/i18n-message-gen/core"
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	cmd := core.NewRootCmd()
	cmd.SetArgs([]string{"--source", "example", "--target", "outputs", "-v"})
	//cmd.SetArgs([]string{"--help"})
	if err := cmd.Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}
