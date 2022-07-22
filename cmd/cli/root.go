/*
Copyright © 2022 Johnson Shi <Johnson.Shi@microsoft.com>
*/
package main

import (
	"flag"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) *cobra.Command {
	opts := &registryRepositoryAuthOpts{}

	cobraCmd := &cobra.Command{
		Use:   "oras-tools",
		Short: "Helper tools for ORAS (https://oras.land)",
	}

	cobraCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	f := cobraCmd.PersistentFlags()

	f.StringVar(&opts.username, "username", "", "username to use for authentication with the registry")
	cobraCmd.MarkFlagRequired("username")

	// TODO add support for --password-stdin (reading password from stdin) for more secure password input.
	f.StringVar(&opts.password, "password", "", "password to use for authentication with the registry")
	cobraCmd.MarkFlagRequired("password")

	f.StringVar(&opts.registry, "registry", "", "hostname of the registry (example: myregistry.azurecr.io)")
	cobraCmd.MarkFlagRequired("registry")

	f.StringVar(&opts.repository, "repository", "", "repository of the artifact within the registry (example: myrepository)")
	cobraCmd.MarkFlagRequired("repository")

	cobraCmd.AddCommand(
		newDeleteCmd(opts, stdin, stdout, stderr, args),
	)

	_ = f.Parse(args)

	return cobraCmd
}

func execute() {
	rootCmd := newRootCmd(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
