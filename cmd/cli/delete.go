/*
Copyright © 2022 Johnson Shi <Johnson.Shi@microsoft.com>
*/
package main

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type deleteOpts struct {
	authOpts *registryRepositoryAuthOpts
	stdin    io.Reader
	stdout   io.Writer
	stderr   io.Writer
	digest   string
}

func newDeleteCmd(authOpts *registryRepositoryAuthOpts, stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) *cobra.Command {
	opts := &deleteOpts{
		authOpts: authOpts,
		stdin:    stdin,
		stdout:   stdout,
		stderr:   stderr,
	}

	cobraCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete ORAS artifacts from a registry",
		RunE: func(_ *cobra.Command, args []string) error {
			return opts.run()
		},
	}

	f := cobraCmd.Flags()

	f.StringVar(&opts.digest, "digest", "", "digest of the artifact (example: sha256:123456789abcdef)")
	cobraCmd.MarkFlagRequired("digest")

	return cobraCmd
}

func (opts *deleteOpts) run() error {
	ctx := context.Background()

	repo, err := opts.authOpts.getAuthenticatedRemoteRepositoryClient()
	if err != nil {
		return err
	}

	ociDescriptor, _, err := repo.FetchReference(ctx, opts.digest)
	if err != nil {
		return err
	}

	err = repo.Delete(ctx, ociDescriptor)
	if err != nil {
		return err
	}

	fmt.Printf("%s/%s@%s deleted successfully\n", opts.authOpts.registry, opts.authOpts.repository, opts.digest)

	return nil
}
