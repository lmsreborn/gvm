package ctl

import (
	"github.com/gvm/command"
	"github.com/spf13/cobra"
)

const (
	cliName        = "gvmctl"
	cliDescription = "A simple command line client for gvm"
)

var (
	rootCmd = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"gvmctl"},
	}
)

func init() {
	rootCmd.AddCommand(
		command.NewVersionCommand(),
	)
}

func init() {
	cobra.EnablePrefixMatching = true
}
